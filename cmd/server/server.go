package main

import (
	authorization "example/movieSuggestion/authorization/server"
	"example/movieSuggestion/database"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
	"example/movieSuggestion/server"
	"example/movieSuggestion/utils"
	"log"
	"net"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port          = ":50501"
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

// This function returns list of roles who can access an endpoint/RPC
func accessibleRoles() map[string][]string {
	const msService = "/msproto.MsDatabase/"
	return map[string][]string{
		msService + "CreateUser":               {"admin"},
		msService + "AddMovie":                 {"admin"},
		msService + "DeleteMovie":              {"admin"},
		msService + "GetAllMovies":             {"admin", "user"},
		msService + "GetMovieByCategory":       {"admin", "user"},
		msService + "AddMovieToWatchlist":      {"admin", "user"},
		msService + "GetAllWatchlistMovies":    {"admin", "user"},
		msService + "DeleteMovieFromWatchlist": {"admin", "user"},
		msService + "CreateReview":             {"admin", "user"},
		msService + "UpdateReview":             {"admin", "user"},
		msService + "CreateLike":               {"admin", "user"},
		msService + "DeleteLike":               {"admin", "user"},
	}
}

func Run(db *gorm.DB) error {
	defer db.Close()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}
	//initialise a JwtManager
	jwtManager := authorization.NewJWTManager(secretKey, tokenDuration)

	//initialise a AuthServer
	authServer := authorization.NewAuthServer(jwtManager, db)

	//initialise a AuthInterceptor
	interceptor := authorization.NewServerAuthInterceptor(jwtManager, accessibleRoles())

	//create new server
	new_server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	//Register the AuthServer Service with the created server
	pb.RegisterAuthServiceServer(new_server, authServer)

	//Register the main service
	pb.RegisterMsDatabaseServer(new_server, &server.MsServer{
		Db: database.DBclient{Db: db},
	})

	reflection.Register(new_server)

	log.Printf("Using port no %v", listen.Addr())

	return new_server.Serve(listen)
}
func main() {
	schema.StartDB()
	// db connection
	db, err := gorm.Open("postgres", utils.GoDotEnvVariable("DB_URL"))
	if err != nil {
		panic(err.Error())
	}

	if err := Run(db); err != nil {
		log.Fatal(err.Error())
	}
}

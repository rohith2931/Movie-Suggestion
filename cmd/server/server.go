package main

import (
	"example/movieSuggestion/authorization"
	"example/movieSuggestion/database"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
	"example/movieSuggestion/server"
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

func accessibleRoles() map[string][]string {
	const msService = "/msproto.MsDatabase/"
	return map[string][]string{
		msService + "CreateUser":  {"admin"},
		msService + "AddMovie":    {"admin"},
		msService + "DeleteMovie": {"admin"},
	}
}
func main() {
	schema.StartDB()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//db connection
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	jwtManager := authorization.NewJWTManager(secretKey, tokenDuration)
	authServer := server.NewAuthServer(jwtManager, db)

	interceptor := authorization.NewAuthInterceptor(jwtManager, accessibleRoles())
	//create new server
	new_server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	pb.RegisterAuthServiceServer(new_server, authServer)
	pb.RegisterMsDatabaseServer(new_server, &server.MsServer{
		Db: database.DBclient{Db: db},
	})
	reflection.Register(new_server)
	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}

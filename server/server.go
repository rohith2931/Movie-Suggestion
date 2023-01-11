package main

// import "example/movieSuggestion/schema"

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":50501"
)

type msServer struct {
	pb.UnimplementedMsDatabaseCrudServer
	Db *gorm.DB
}

func (s *msServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("creating new user called")
	newUser := schema.User{
		UserName:    in.GetUserName(),
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
		Address:     in.GetAddress(),
		Watchlist:   schema.Watchlist{Count: 0},
	}
	s.Db.Create(&newUser)
	return &pb.User{UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber(), Id: uint64(newUser.ID)}, nil
}

func (s *msServer) GetAllMovies(in *pb.EmptyMovie, stream pb.MsDatabaseCrud_GetAllMoviesServer) error {
	log.Printf("Getting movies called")
	Movies := []schema.Movie{}
	// AllMovies := []*pb.Movie{}
	s.Db.Find(&Movies)
	for _, movie := range Movies {
		Movie := &pb.Movie{
			Id:          uint64(movie.ID),
			Name:        movie.Name,
			Director:    movie.Director,
			Description: movie.Description,
			Rating:      uint64(movie.Rating),
			Language:    movie.Language,
			Category:    movie.Category,
			ReleaseDate: movie.ReleaseDate,
		}
		if err := stream.Send(Movie); err != nil {
			return err
		}
	}
	return nil
}

func (s *msServer) GetMovieByCategory(ctx context.Context, in *pb.MovieCategory) (*pb.Movies, error) {
	log.Printf("get by category is called")
	Movies := []schema.Movie{}
	AllMovies := []*pb.Movie{}
	s.Db.Model(&schema.Movie{}).Where("Category=?", in.Category).Find(&Movies)
	for _, movie := range Movies {
		AllMovies = append(AllMovies, &pb.Movie{
			Id:          uint64(movie.ID),
			Name:        movie.Name,
			Director:    movie.Director,
			Description: movie.Description,
			Rating:      uint64(movie.Rating),
			Language:    movie.Language,
			Category:    movie.Category,
			ReleaseDate: movie.ReleaseDate,
		})
	}
	return &pb.Movies{Movies: AllMovies}, nil
}
func (s *msServer) AddMovie(ctx context.Context, in *pb.NewMovie) (*pb.Movie, error) {
	log.Printf("Add Movie is called")
	newMovie := schema.Movie{
		Name:        in.GetName(),
		Director:    in.GetDirector(),
		Description: in.GetDescription(),
		Rating:      int(in.GetRating()),
		Language:    in.GetLanguage(),
		Category:    in.GetCategory(),
		ReleaseDate: in.GetReleaseDate(),
	}
	s.Db.Save(&newMovie)
	return &pb.Movie{
		Name:        newMovie.Name,
		Director:    newMovie.Director,
		Description: newMovie.Description,
		Rating:      uint64(newMovie.Rating),
		Language:    newMovie.Language,
		Category:    newMovie.Category,
		ReleaseDate: newMovie.ReleaseDate,
		Id:          uint64(newMovie.ID)}, nil
}
func (s *msServer) DeleteMovie(ctx context.Context, in *pb.Movie) (*pb.Movie, error) {
	log.Printf("Delete Movie is called")
	movie := schema.Movie{}
	s.Db.Find(&movie, in.Id)
	s.Db.Delete(&schema.Movie{}, in.Id)
	return &pb.Movie{
		Id:          uint64(movie.ID),
		Name:        movie.Name,
		Director:    movie.Director,
		Description: movie.Description,
		Rating:      uint64(movie.Rating),
		Language:    movie.Language,
		Category:    movie.Category,
		ReleaseDate: movie.ReleaseDate,
	}, nil
}
func (s *msServer) AddMovieToWatchlist(ctx context.Context, in *pb.AddMovieByUser) (*pb.Movie, error) {
	fmt.Println("Add to Movie by user is called")
	user := schema.User{}
	s.Db.Preload("Watchlist").Find(&user, in.UserId)
	movie := schema.Movie{}
	s.Db.Find(&movie, in.MovieId)
	fmt.Println(user.ID, movie.ID)
	// fmt.Println(user.Watchlist)
	// fmt.Println(user)
	user.Watchlist.Movies = append(user.Watchlist.Movies, movie)
	user.Watchlist.Count = len(user.Watchlist.Movies)
	s.Db.Save(&user)
	fmt.Println()
	// fmt.Println(user.Watchlist)
	// fmt.Println(user)

	return &pb.Movie{}, nil
}
func (s *msServer) CreateReview(ctx context.Context, in *pb.NewReview) (*pb.Review, error) {
	newReview := schema.Review{
		Rating:      int(in.Rating),
		MovieID:     uint(in.MovieId),
		UserID:      uint(in.UserId),
		Description: in.Description,
	}
	s.Db.Create(&newReview)
	return &pb.Review{
		Id:          uint64(newReview.ID),
		Rating:      uint64(newReview.Rating),
		MovieId:     uint64(newReview.MovieID),
		UserId:      uint64(newReview.UserID),
		Description: newReview.Description,
	}, nil
}

func (s *msServer) UpdateReview(ctx context.Context, in *pb.Review) (*pb.Review, error) {
	log.Printf("Update Review is called")
	review := schema.Review{
		Rating:      int(in.Rating),
		MovieID:     uint(in.MovieId),
		UserID:      uint(in.UserId),
		Description: in.Description,
	}
	s.Db.Model(&schema.Review{}).Where("id=?", in.Id).Updates(review)
	return &pb.Review{
		Id:          uint64(in.Id),
		UserId:      uint64(review.UserID),
		MovieId:     uint64(review.MovieID),
		Rating:      uint64(review.Rating),
		Description: review.Description,
	}, nil
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

	//create new server
	new_server := grpc.NewServer()
	pb.RegisterMsDatabaseCrudServer(new_server, &msServer{
		Db: db,
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}

package main

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50501" // port address for client
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	schema.CheckError(err)
	defer conn.Close()

	client := pb.NewMsDatabaseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Creating a new user
	new_user, err := client.CreateUser(ctx, &pb.NewUser{
		UserName:    "sam",
		Password:    "sam",
		Email:       "sam@mail",
		PhoneNumber: "7454532421",
		Address:     "LA",
	})
	schema.CheckError(err)
	log.Printf("\nId : %v,User Name: %v, Password: %v, Email: %v, PhoneNumber : %v ", new_user.Id, new_user.GetUserName(), new_user.GetPassword(), new_user.GetEmail(), new_user.GetPhoneNumber())

	//Get All Movies
	stream, err := client.GetAllMovies(ctx, &pb.EmptyMovie{})
	schema.CheckError(err)
	fmt.Println("All movies:")
	for {
		movie, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Printf("%v\n", movie)
	}

	// for _, movie := range AllMovies.Movies {
	// 	fmt.Println(movie.GetId(), movie.GetName(), movie.GetDirector(), movie.GetDescription(), movie.GetRating(), movie.GetLanguage(), movie.GetCategory(), movie.GetReleaseDate())
	// }

	//Create a movie
	NewMovie, err := client.AddMovie(ctx, &pb.NewMovie{
		Name:        "Avengers",
		Director:    "Joss Whedon",
		Description: "Superhero Film",
		Rating:      8,
		Language:    "English",
		Category:    "Action",
		ReleaseDate: "04-05-2012",
	})
	schema.CheckError(err)
	fmt.Println("Created movie")
	fmt.Println(NewMovie.GetId(), NewMovie.GetName(), NewMovie.GetDirector(), NewMovie.GetDescription(), NewMovie.GetRating(), NewMovie.GetLanguage(), NewMovie.GetCategory(), NewMovie.GetReleaseDate())

	//Get all Movies by category
	AllMoviesByCategory, err := client.GetMovieByCategory(ctx, &pb.MovieCategory{Category: "Drama"})
	schema.CheckError(err)
	log.Println("search by Category")
	for _, movie := range AllMoviesByCategory.Movies {
		fmt.Println(movie.GetId(), movie.GetName(), movie.GetDirector(), movie.GetDescription(), movie.GetRating(), movie.GetLanguage(), movie.GetCategory(), movie.GetReleaseDate())
	}

	// Add movie to Watchlist by user
	fmt.Println("Adding movie to watchlist:")
	Movie, err := client.AddMovieToWatchlist(ctx, &pb.AddMovieByUser{
		UserId:  2,
		MovieId: 2,
	})
	schema.CheckError(err)
	// fmt.Printf("%+v\n", Movie)
	fmt.Println(Movie.GetId(), Movie.GetName(), Movie.GetDirector(), Movie.GetDescription(), Movie.GetRating(), Movie.GetLanguage(), Movie.GetCategory(), Movie.GetReleaseDate())

	Movie, err = client.AddMovieToWatchlist(ctx, &pb.AddMovieByUser{
		UserId:  2,
		MovieId: 1,
	})
	schema.CheckError(err)
	// fmt.Printf("%+v\n", Movie)
	fmt.Println(Movie.GetId(), Movie.GetName(), Movie.GetDirector(), Movie.GetDescription(), Movie.GetRating(), Movie.GetLanguage(), Movie.GetCategory(), Movie.GetReleaseDate())

	//Delete a movie
	// DeletedMovie, err := client.DeleteMovie(ctx, &pb.Movie{
	// 	Id: 1,
	// })
	// schema.CheckError(err)
	// fmt.Println("Deleted movie is:")
	// fmt.Println(DeletedMovie.GetId(), DeletedMovie.GetName(), DeletedMovie.GetDirector(), DeletedMovie.GetDescription(), DeletedMovie.GetRating(), NewMovie.GetLanguage(), DeletedMovie.GetCategory(), DeletedMovie.GetReleaseDate())

	//Create Review to a movie
	newReview, err := client.CreateReview(ctx, &pb.NewReview{
		Rating:      8,
		MovieId:     2,
		UserId:      2,
		Description: "excellent",
	})
	schema.CheckError(err)
	fmt.Printf("\nCreated review is:\n %v\n", newReview)

	//Update Review to a movie
	updatedReview, err := client.UpdateReview(ctx, &pb.Review{
		Id:          1,
		Rating:      5,
		MovieId:     1,
		UserId:      1,
		Description: "Very Good",
	})
	schema.CheckError(err)
	fmt.Printf("\nUpdated review is:\n %v\n", updatedReview)

	//Get all movies from the watchlist of a user
	watchlistMovies, err := client.GetAllWatchlistMovies(ctx, &pb.UserId{Id: 2})
	schema.CheckError(err)
	fmt.Println("Watchlist Movies of the User is :")
	for _, movie := range watchlistMovies.Movies {
		fmt.Printf("%v\n", movie)
	}

	fmt.Println("Deleting movie from watchlist:")
	DeletedMovieByUser, err := client.DeleteMovieFromWatchlist(ctx, &pb.DeleteMovieByUser{
		UserId:  2,
		MovieId: 2,
	})
	schema.CheckError(err)
	fmt.Printf("%+v\n", DeletedMovieByUser)

	//Creating a like by a user
	RespnseObj, err := client.CreateLike(ctx, &pb.UserLike{
		UserId:  1,
		MovieId: 1,
	})
	schema.CheckError(err)
	fmt.Println(RespnseObj.Body)

	RespnseObj, err = client.CreateLike(ctx, &pb.UserLike{
		UserId:  1,
		MovieId: 2,
	})
	schema.CheckError(err)
	fmt.Println(RespnseObj.Body)

	//Deleting a like by a user
	RespnseObj, err = client.DeleteLike(ctx, &pb.UserLike{
		UserId:  1,
		MovieId: 2,
	})
	schema.CheckError(err)
	fmt.Println(RespnseObj.Body)

}

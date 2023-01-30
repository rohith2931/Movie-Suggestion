package main

import (
	"context"
	"example/movieSuggestion/authorization/client"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/utils"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address         = "localhost:50501" // port address for client
	username        = "rohith"
	password        = "pass"
	refreshDuration = 30 * time.Second
)

// This function returns which RPC/endpoint needs authentication
func accessibleMethods() map[string]bool {
	const msService = "/msproto.MovieSuggestionService/"
	return map[string]bool{
		msService + "CreateUser":               true,
		msService + "AddMovie":                 true,
		msService + "DeleteMovie":              true,
		msService + "GetAllMovies":             true,
		msService + "GetMovieByCategory":       true,
		msService + "AddMovieToWatchlist":      true,
		msService + "GetAllWatchlistMovies":    true,
		msService + "DeleteMovieFromWatchlist": true,
		msService + "CreateReview":             true,
		msService + "UpdateReview":             true,
		msService + "CreateLike":               true,
		msService + "DeleteLike":               true,
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	utils.PanicError(err)
	defer conn.Close()

	//Create a new AuthClient with the provided Credentials
	authClient := client.NewAuthClient(conn, username, password)

	//Create client AuthInterceptor
	interceptor, err := client.NewClientAuthInterceptor(authClient, accessibleMethods(), refreshDuration)
	if err != nil {
		log.Println("cannot create auth interceptor:", err)
	}

	clientConnection, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	utils.PanicError(err)
	defer clientConnection.Close()
	client := pb.NewMovieSuggestionServiceClient(clientConnection)

	// createNewUser(client)
	createNewMovie(client)
	// deleteMovie(client)
	// addMovieToWatchlist(client)
	// getAllMovies(client)
	// getMoviesByCategory(client)
	// createReview(client)
	// updateReview(client)
	// deleteMovieFromWatchlist(client)
	// getAllWatchlistMovies(client)
	// createLike(client)
	// deleteLike(client)

}

// func createNewUser(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	//Creating a new user
// 	new_user, err := client.CreateUser(ctx, &pb.NewUser{
// 		UserName:    "sam",
// 		Password:    "sam",
// 		Email:       "seropa3842@quamox.com",
// 		PhoneNumber: "7454532421",
// 		Address:     "LA",
// 	})
// 	utils.CheckError(err)
// 	log.Printf("\nId : %v,User Name: %v, Password: %v, Email: %v, PhoneNumber : %v ", new_user.Id, new_user.GetUserName(), new_user.GetPassword(), new_user.GetEmail(), new_user.GetPhoneNumber())
// }

// func addMovieToWatchlist(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	// Add a movie to Watchlist by user
// 	fmt.Println("Adding movie to watchlist:")
// 	Movie, err := client.AddMovieToWatchlist(ctx, &pb.AddMovieByUser{
// 		UserId:  1,
// 		MovieId: 1,
// 	})
// 	utils.CheckError(err)
// 	fmt.Println(Movie.GetId(), Movie.GetName(), Movie.GetDirector(), Movie.GetDescription(), Movie.GetRating(), Movie.GetLanguage(), Movie.GetCategory(), Movie.GetReleaseDate())
// }

// func getAllWatchlistMovies(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	// Get all movies from the watchlist of a user
// 	watchlistMovies, err := client.GetAllWatchlistMovies(ctx, &pb.UserId{Id: 1})
// 	utils.CheckError(err)
// 	fmt.Println("Watchlist Movies of the User is :")
// 	for _, movie := range watchlistMovies.Movies {
// 		fmt.Printf("%v\n", movie)
// 	}
// }

// func deleteMovieFromWatchlist(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	// Delete a Movie from watchlist of a user
// 	fmt.Println("Deleting movie from watchlist:")
// 	DeletedMovieByUser, err := client.DeleteMovieFromWatchlist(ctx, &pb.DeleteMovieByUser{
// 		UserId:  1,
// 		MovieId: 4,
// 	})
// 	utils.CheckError(err)
// 	fmt.Printf("%+v\n", DeletedMovieByUser)
// }

// func createLike(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	// Create a like by a user to a movie
// 	RespnseObj, err := client.CreateLike(ctx, &pb.UserLike{
// 		UserId:  1,
// 		MovieId: 2,
// 	})
// 	utils.CheckError(err)
// 	fmt.Println(RespnseObj.Body)
// }

// func deleteLike(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	// Delete/Unlike a movie by a user
// 	RespnseObj, err := client.DeleteLike(ctx, &pb.UserLike{
// 		UserId:  1,
// 		MovieId: 2,
// 	})
// 	utils.CheckError(err)
// 	fmt.Println(RespnseObj.Body)
// }

func createNewMovie(client pb.MovieSuggestionServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Create a movie into the database
	NewMovie, err := client.AddMovie(ctx, &pb.Movie{
		Name:        "Avengers",
		Director:    "Joss Whedon",
		Description: "Superhero Film",
		// Rating:      8,
		Language:    "English",
		Category:    "Action",
		ReleaseDate: "04-05-2012",
	})
	utils.CheckError(err)
	fmt.Println("Created movie")
	fmt.Println(NewMovie.GetMovieId(), NewMovie.GetName(), NewMovie.GetDirector(), NewMovie.GetDescription(), NewMovie.GetRating(), NewMovie.GetLanguage(), NewMovie.GetCategory(), NewMovie.GetReleaseDate())
}

// func deleteMovie(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Delete a movie in the database
// 	DeletedMovie, err := client.DeleteMovie(ctx, &pb.Movie{
// 		Id: 1,
// 	})
// 	utils.CheckError(err)
// 	fmt.Println("Deleted movie is:")
// 	fmt.Println(DeletedMovie.GetId(), DeletedMovie.GetName(), DeletedMovie.GetDirector(), DeletedMovie.GetDescription(), DeletedMovie.GetRating(), DeletedMovie.GetLanguage(), DeletedMovie.GetCategory(), DeletedMovie.GetReleaseDate())
// }

// func getAllMovies(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	//Get All Movies that can be watched
// 	stream, err := client.GetAllMovies(ctx, &pb.EmptyMovie{})
// 	utils.CheckError(err)
// 	fmt.Println("All movies: ")
// 	for {
// 		movie, err := stream.Recv()
// 		if err == io.EOF || movie == nil {
// 			break
// 		}
// 		fmt.Printf("%v\n", movie)
// 	}
// }

// func getMoviesByCategory(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	//Get all Movies based on category/genre
// 	AllMoviesByCategory, err := client.GetMovieByCategory(ctx, &pb.MovieCategory{Category: "Action"})
// 	utils.CheckError(err)
// 	fmt.Println("search by Category")
// 	for _, movie := range AllMoviesByCategory.Movies {
// 		fmt.Println(movie.GetId(), movie.GetName(), movie.GetDirector(), movie.GetDescription(), movie.GetRating(), movie.GetLanguage(), movie.GetCategory(), movie.GetReleaseDate())
// 	}
// }

// func createReview(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	//Create a Review to a movie
// 	newReview, err := client.CreateReview(ctx, &pb.NewReview{
// 		Rating:      8,
// 		MovieId:     2,
// 		UserId:      2,
// 		Description: "excellent",
// 	})
// 	utils.CheckError(err)
// 	fmt.Printf("\nCreated review is:\n %v\n", newReview)

// }
// func updateReview(client pb.MovieSuggestionServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	//Update Review to a movie
// 	updatedReview, err := client.UpdateReview(ctx, &pb.Review{
// 		Id:          1,
// 		Rating:      5,
// 		MovieId:     1,
// 		UserId:      1,
// 		Description: "Very Good",
// 	})
// 	utils.CheckError(err)
// 	fmt.Printf("\nUpdated review is:\n %v\n", updatedReview)

// }

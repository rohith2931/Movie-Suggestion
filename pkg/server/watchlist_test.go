package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/pkg/database"
	"example/movieSuggestion/pkg/schema"
	"example/movieSuggestion/utils"
	"fmt"
	"io"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestAddMovieToWatchlist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	expected := &pb.AddMovieToWatchlistResponse{
		UserId:  1,
		MovieId: 1,
		Status:  "Movie Added Successfully",
	}
	mockDb.EXPECT().AddMovieToWatchlist(gomock.Any(), gomock.Any()).Return(nil)
	got, err := msServer.AddMovieToWatchlist(ctx, &pb.AddMovieToWatchlistRequest{
		MovieId: 1,
		UserId:  1,
	})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestGetAllWatchlistMovies(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	c, closer := MockServer(mockDb, ctx)
	defer closer()
	mockDb.EXPECT().GetAllWatchlistMovies(gomock.Any()).Return(
		[]schema.Movie{
			{
				Name:        "F2",
				Director:    "Anil Ravipudi",
				Description: "Fun and frustation",
				Rating:      7,
				Language:    "Telugu",
				Category:    "Drama",
				ReleaseDate: "12-01-2019",
			},
		}, nil)
	want := []*pb.Movie{
		{
			Name:        "F2",
			Director:    "Anil Ravipudi",
			Description: "Fun and frustation",
			Rating:      7,
			Language:    "Telugu",
			Category:    "Drama",
			ReleaseDate: "12-01-2019",
		},
	}
	stream, err := c.GetAllWatchlistMovies(ctx, &pb.GetAllWatchlistMoviesRequest{
		UserId: 1,
	})
	utils.CheckError(err)
	for _, expected := range want {
		got, err := stream.Recv()

		if err == io.EOF {
			log.Printf("finished")
			return
		}
		utils.CheckError(err)
		fmt.Printf("expected %v\n", expected)
		fmt.Println(reflect.DeepEqual(got, expected))
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("The Function Retured is not expected one.\n got %v \n expected %v",
				got, expected)
		}
	}
}

func TestDeleteMovieFromWatchlist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	expected := &pb.DeleteMovieFromWatchlistResponse{
		UserId:  1,
		MovieId: 1,
		Status:  "Movie Deleted from watchlist successfully",
	}
	mockDb.EXPECT().DeleteMovieFromWatchlist(gomock.Any(), gomock.Any()).Return(nil)
	got, err := msServer.DeleteMovieFromWatchlist(ctx, &pb.DeleteMovieFromWatchlistRequest{
		MovieId: 1,
		UserId:  1,
	})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}
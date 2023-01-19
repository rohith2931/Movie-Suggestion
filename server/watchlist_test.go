package server

import (
	"context"
	"example/movieSuggestion/database"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/utils"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddMovieToWatchlist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	expected := &pb.Movie{
		Name:        "F2",
		Director:    "Anil Ravipudi",
		Description: "Fun and frustation",
		Rating:      7,
		Language:    "Telugu",
		Category:    "Drama",
		ReleaseDate: "12-01-2019",
	}
	mockDb.EXPECT().AddMovieToWatchlist(gomock.Any(), gomock.Any()).Return(sampleMovie)
	got, err := msServer.AddMovieToWatchlist(ctx, &pb.AddMovieByUser{
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
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	expected := []*pb.Movie{
		{
			Name:        "F2",
			Director:    "Anil Ravipudi",
			Description: "Fun and frustation",
			Rating:      7,
			Language:    "Telugu",
			Category:    "Drama",
			ReleaseDate: "12-01-2019",
		},
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
	mockDb.EXPECT().GetAllWatchlistMovies(gomock.Any()).Return(sampleMovies)
	got, err := msServer.GetAllWatchlistMovies(ctx, &pb.UserId{
		Id: 1,
	})
	utils.CheckError(err)
	if !reflect.DeepEqual(got.Movies, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got.Movies, expected)
	}
}

func TestDeleteMovieFromWatchlist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	expected := &pb.Movie{
		Name:        "F2",
		Director:    "Anil Ravipudi",
		Description: "Fun and frustation",
		Rating:      7,
		Language:    "Telugu",
		Category:    "Drama",
		ReleaseDate: "12-01-2019",
	}
	mockDb.EXPECT().DeleteMovieFromWatchlist(gomock.Any(), gomock.Any()).Return(sampleMovie)
	got, err := msServer.DeleteMovieFromWatchlist(ctx, &pb.DeleteMovieByUser{
		MovieId: 1,
		UserId:  1,
	})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

package server

import (
	"context"
	"example/movieSuggestion/database"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
	"example/movieSuggestion/utils"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"google.golang.org/grpc"
)

var sampleMovies = []schema.Movie{

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

var sampleMovie = schema.Movie{
	Name:        "F2",
	Director:    "Anil Ravipudi",
	Description: "Fun and frustation",
	Rating:      7,
	Language:    "Telugu",
	Category:    "Drama",
	ReleaseDate: "12-01-2019",
}

type mockMsDatabase_GetAllMoviesServer struct {
	grpc.ServerStream
	Results []*pb.Movie
}

func NewmockMsDatabase_GetAllMoviesServer() *mockMsDatabase_GetAllMoviesServer {
	return &mockMsDatabase_GetAllMoviesServer{
		Results: make([]*pb.Movie, 0),
	}
}
func (mock mockMsDatabase_GetAllMoviesServer) Send(m *pb.Movie) error {
	temp := &mock.Results
	mock.Results = append(*temp, m)
	// mock.Results = append(mock.Results, m)
	return nil
}
func TestGetAllMovies(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	mockDb.EXPECT().GetAllMovies().Return(sampleMovies)
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
	mock := NewmockMsDatabase_GetAllMoviesServer()
	err := msServer.GetAllMovies(&pb.EmptyMovie{}, mock)
	utils.CheckError(err)
	if !reflect.DeepEqual(mock.Results, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			mock.Results, expected)
	}
}
func TestGetMovieByCategory(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().GetMovieByCategory(gomock.Any()).Return(sampleMovies)
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
	got, err := msServer.GetMovieByCategory(ctx, &pb.MovieCategory{
		Category: "Drama",
	})
	utils.CheckError(err)
	if !reflect.DeepEqual(got.Movies, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got.Movies, expected)
	}
}

func TestAddMovie(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddMovie(gomock.Any())
	expected := &pb.Movie{
		Name:        "F2",
		Director:    "Anil Ravipudi",
		Description: "Fun and frustation",
		Rating:      7,
		Language:    "Telugu",
		Category:    "Drama",
		ReleaseDate: "12-01-2019",
	}
	got, err := msServer.AddMovie(ctx, &pb.NewMovie{
		Name:        "F2",
		Director:    "Anil Ravipudi",
		Description: "Fun and frustation",
		Rating:      7,
		Language:    "Telugu",
		Category:    "Drama",
		ReleaseDate: "12-01-2019"})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestDeleteMovie(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().DeleteMovie(gomock.Any()).Return(sampleMovie)
	expected := &pb.Movie{
		Name:        "F2",
		Director:    "Anil Ravipudi",
		Description: "Fun and frustation",
		Rating:      7,
		Language:    "Telugu",
		Category:    "Drama",
		ReleaseDate: "12-01-2019",
	}
	got, err := msServer.DeleteMovie(ctx, &pb.Movie{Id: 1})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

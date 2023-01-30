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

var sampleMovie = schema.Movie{
	Name:        "F2",
	Director:    "Anil Ravipudi",
	Description: "Fun and frustation",
	Rating:      7,
	Language:    "Telugu",
	Category:    "Drama",
	ReleaseDate: "12-01-2019",
}

func TestGetAllMovies(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	c, closer := MockServer(mockDb, ctx)
	defer closer()
	mockDb.EXPECT().GetAllMovies(gomock.Any(), gomock.Any()).Return([]schema.Movie{
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
	stream, err := c.GetAllMovies(ctx, &pb.MovieRequest{})
	utils.CheckError(err)
	// got, err := stream.Recv()

	// if err == io.EOF {
	// 	log.Printf("finished")
	// 	return
	// }
	// fmt.Println(got)
	// fmt.Println(want[0])
	// fmt.Println(reflect.DeepEqual(got, want[0]))
	for _, expected := range want {
		got, err := stream.Recv()

		if err == io.EOF {
			log.Printf("finished")
			return
		}
		utils.CheckError(err)
		// fmt.Printf("Got %v\n", got)
		// fmt.Printf("%+v\n", expected)
		// time.Sleep(time.Second * 2)
		// fmt.Printf("%+v\n", got)
		fmt.Printf("expected %v\n", expected)
		// fmt.Println("hii")
		// fmt.Println("is equal", []byte(fmt.Sprintf("%v", got)), []byte(fmt.Sprintf("%v", expected)))
		fmt.Println(reflect.DeepEqual(got, expected))
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("The Function Retured is not expected one.\n got %v \n expected %v",
				got, expected)
		}
		// 	// fmt.Printf("%+v\n", got)
		// 	// fmt.Printf("%+v\n", expected)

	}
}

func TestAddMovie(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddMovie(gomock.Any()).Return(nil)
	expected := &pb.Movie{
		Name:        "F2",
		Director:    "Anil Ravipudi",
		Description: "Fun and frustation",
		Rating:      7,
		Language:    "Telugu",
		Category:    "Drama",
		ReleaseDate: "12-01-2019",
	}
	got, err := msServer.AddMovie(ctx, &pb.Movie{
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
	mockDb.EXPECT().DeleteMovie(gomock.Any()).Return(nil)
	expected := &pb.DeleteMovieResponse{
		MovieId: 1,
		Status:  "Movie Deleted Successfully",
	}
	got, err := msServer.DeleteMovie(ctx, &pb.Movie{MovieId: 1})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

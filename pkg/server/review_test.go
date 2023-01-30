package server

import (
	"context"
	"example/movieSuggestion/pkg/database"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/pkg/schema"
	"example/movieSuggestion/utils"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var sampleReview = schema.Review{
	Rating:      8,
	MovieID:     1,
	UserID:      1,
	Description: "Good",
}

func TestCreateReview(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()

	mockDb.EXPECT().CreateReview(gomock.Any()).Return(nil)
	expected := &pb.Review{
		Rating:      8,
		MovieId:     1,
		UserId:      1,
		Description: "Good",
	}
	got, err := msServer.CreateReview(ctx, &pb.Review{
		Rating:      8,
		MovieId:     1,
		UserId:      1,
		Description: "Good",
	})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestUpdateReview(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()

	mockDb.EXPECT().UpdateReview(gomock.Any(), gomock.Any()).Return(nil)
	expected := &pb.Review{
		Rating:      8,
		MovieId:     1,
		UserId:      1,
		Description: "Good",
	}
	got, err := msServer.UpdateReview(ctx, &pb.Review{
		Rating:      8,
		MovieId:     1,
		UserId:      1,
		Description: "Good",
	})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

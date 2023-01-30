package server

import (
	"context"
	"reflect"
	"testing"

	"example/movieSuggestion/pkg/database"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/utils"

	"github.com/golang/mock/gomock"
)

func TestCreateLike(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().CreateLike(gomock.Any(), gomock.Any()).Return("Like created Successfully..", nil)
	expected := pb.UserLikeResponse{
		Status: "Like created Successfully..",
	}
	got, err := msServer.CreateLike(ctx, &pb.UserLikeRequest{MovieId: 1, UserId: 1})
	utils.CheckError(err)
	if !reflect.DeepEqual(got.Status, expected.Status) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got.Status, expected.Status)
	}
}

func TestDeleteLike(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().DeleteLike(gomock.Any(), gomock.Any()).Return("Like Deleted Successfully..", nil)
	expected := pb.UserLikeResponse{
		Status: "Like Deleted Successfully..",
	}
	got, err := msServer.DeleteLike(ctx, &pb.UserLikeRequest{MovieId: 1, UserId: 1})
	utils.CheckError(err)
	if !reflect.DeepEqual(got.Status, expected.Status) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got.Status, expected.Status)
	}
}

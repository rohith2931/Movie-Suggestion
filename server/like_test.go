package server

import (
	"context"
	"reflect"
	"testing"

	"example/movieSuggestion/mocks"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"

	"github.com/golang/mock/gomock"
)

func TestCreateLike(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().CreateLike(gomock.Any(), gomock.Any()).Return("Like created Successfully..")
	expected := pb.Response{
		Body: "Like created Successfully..",
	}
	got, err := msServer.CreateLike(ctx, &pb.UserLike{MovieId: 1, UserId: 1})
	schema.CheckError(err)
	if !reflect.DeepEqual(got.Body, expected.Body) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got.Body, expected.Body)
	}
}

func TestDeleteLike(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().DeleteLike(gomock.Any(), gomock.Any()).Return("Like Deleted Successfully..")
	expected := pb.Response{
		Body: "Like Deleted Successfully..",
	}
	got, err := msServer.DeleteLike(ctx, &pb.UserLike{MovieId: 1, UserId: 1})
	schema.CheckError(err)
	if !reflect.DeepEqual(got.Body, expected.Body) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got.Body, expected.Body)
	}
}

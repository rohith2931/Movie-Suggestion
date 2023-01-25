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

var sampleUser = schema.User{
	UserName:    "ram",
	Password:    "passss",
	Email:       "ram@gmail.com",
	PhoneNumber: "8455612342",
	Address:     "Hyderabad",
	Role:        "user",
	Watchlist:   &schema.Watchlist{},
}

func TestCreateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	msServer := MsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().CreateUser(&sampleUser).Return(nil)
	expected := &pb.User{
		UserName:    "ram",
		Password:    "passss",
		Email:       "ram@gmail.com",
		PhoneNumber: "8455612342",
		Address:     "Hyderabad",
	}
	got, err := msServer.CreateUser(ctx, &pb.NewUser{
		UserName:    "ram",
		Password:    "passss",
		Email:       "ram@gmail.com",
		PhoneNumber: "8455612342",
		Address:     "Hyderabad",
	})
	utils.CheckError(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

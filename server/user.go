package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
)

func (s *MsServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	newUser := schema.User{
		UserName:    in.GetUserName(),
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
		Address:     in.GetAddress(),
		Role:        "user",
		Watchlist:   &schema.Watchlist{},
	}
	s.Db.CreateUser(&newUser)
	return &pb.User{UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber(), Id: uint64(newUser.ID), Address: in.Address}, nil
}

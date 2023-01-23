package server

import (
	"context"
	"example/movieSuggestion/Mail"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
)

// This RPC creates a user
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
	mailInfo := Mail.MailInfo(in.GetUserName(), in.GetEmail())
	go Mail.SendMail(in.GetUserName(), in.GetEmail(), mailInfo)
	return &pb.User{UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber(), Id: uint64(newUser.ID), Address: in.Address}, nil
}

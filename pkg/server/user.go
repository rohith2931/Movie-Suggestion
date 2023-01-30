package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/pkg/Mail"
	"example/movieSuggestion/pkg/schema"
	"log"
)

// This RPC creates a user
func (s *MsServer) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	newUser := schema.User{
		UserName:    in.GetUserName(),
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
		Address:     in.GetAddress(),
		Role:        "user",
		Watchlist:   &schema.Watchlist{},
	}
	err := s.Db.CreateUser(&newUser)
	if err != nil {
		log.Println(err)
		return &pb.User{}, err
	}
	mailInfo := Mail.MailInfo(in.GetUserName(), in.GetEmail())
	go Mail.SendMail(mailInfo)
	return &pb.User{UserId: uint64(newUser.ID), UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber(), Address: in.Address}, nil
}

package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/pkg/Mail"
	"example/movieSuggestion/pkg/schema"
	"example/movieSuggestion/utils"
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
	err := s.Db.CreateUser(&newUser)
	utils.CheckError(err)
	mailInfo := Mail.MailInfo(in.GetUserName(), in.GetEmail())
	go Mail.SendMail(mailInfo)
	return &pb.User{UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber(), Id: uint64(newUser.ID), Address: in.Address}, nil
}

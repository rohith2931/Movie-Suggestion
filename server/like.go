package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
)

func (s *MsServer) CreateLike(ctx context.Context, in *pb.UserLike) (*pb.Response, error) {
	status := s.Db.CreateLike(in.UserId, in.MovieId)
	return &pb.Response{Body: status}, nil
}

func (s *MsServer) DeleteLike(ctx context.Context, in *pb.UserLike) (*pb.Response, error) {
	status := s.Db.DeleteLike(in.UserId, in.MovieId)
	return &pb.Response{Body: status}, nil
}

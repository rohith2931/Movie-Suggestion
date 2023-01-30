package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"log"
)

// This RPC Creates a like for a movie for a user
func (s *MsServer) CreateLike(ctx context.Context, in *pb.UserLikeRequest) (*pb.UserLikeResponse, error) {
	status, err := s.Db.CreateLike(in.UserId, in.MovieId)
	if err != nil {
		log.Println(status)
		return &pb.UserLikeResponse{
			UserId:  in.GetUserId(),
			MovieId: in.GetMovieId(),
			Status:  status,
		}, err
	}
	return &pb.UserLikeResponse{
		UserId:  in.GetUserId(),
		MovieId: in.GetMovieId(),
		Status:  status,
	}, nil
}

// This RPC deletes the like for a movie for a user
func (s *MsServer) DeleteLike(ctx context.Context, in *pb.UserLikeRequest) (*pb.UserLikeResponse, error) {
	status, err := s.Db.DeleteLike(in.UserId, in.MovieId)
	if err != nil {
		log.Println(status)
		return &pb.UserLikeResponse{
			UserId:  in.GetUserId(),
			MovieId: in.GetMovieId(),
			Status:  status,
		}, err
	}
	return &pb.UserLikeResponse{
		UserId:  in.GetUserId(),
		MovieId: in.GetMovieId(),
		Status:  status,
	}, nil
}

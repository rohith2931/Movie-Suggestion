package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
)

func (s *MsServer) CreateLike(ctx context.Context, in *pb.UserLike) (*pb.Response, error) {
	s.Db.Create(&schema.Like{
		UserID:  uint(in.UserId),
		MovieID: uint(in.MovieId),
	})
	return &pb.Response{Body: "Like created Successfully.."}, nil
}

func (s *MsServer) DeleteLike(ctx context.Context, in *pb.UserLike) (*pb.Response, error) {
	s.Db.Unscoped().Model(&schema.Like{}).Where("user_id=? and movie_id=?", in.UserId, in.MovieId).Delete(&schema.Like{})
	return &pb.Response{Body: "Like Deleted Successfully.."}, nil
}

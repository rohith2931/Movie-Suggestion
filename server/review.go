package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
)

// This RPC creates a review for a movie by the user
func (s *MsServer) CreateReview(ctx context.Context, in *pb.NewReview) (*pb.Review, error) {
	newReview := schema.Review{
		Rating:      int(in.Rating),
		MovieID:     uint(in.MovieId),
		UserID:      uint(in.UserId),
		Description: in.Description,
	}
	s.Db.CreateReview(&newReview)
	return &pb.Review{
		Id:          uint64(newReview.ID),
		Rating:      uint64(newReview.Rating),
		MovieId:     uint64(newReview.MovieID),
		UserId:      uint64(newReview.UserID),
		Description: newReview.Description,
	}, nil
}

// This RPC updates a review for a movie by the user
func (s *MsServer) UpdateReview(ctx context.Context, in *pb.Review) (*pb.Review, error) {
	review := schema.Review{
		Rating:      int(in.Rating),
		MovieID:     uint(in.MovieId),
		UserID:      uint(in.UserId),
		Description: in.Description,
	}
	s.Db.UpdateReview(in.Id, review)
	return &pb.Review{
		Id:          uint64(in.Id),
		UserId:      uint64(review.UserID),
		MovieId:     uint64(review.MovieID),
		Rating:      uint64(review.Rating),
		Description: review.Description,
	}, nil
}

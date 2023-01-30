package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/pkg/schema"
	"log"
)

// This RPC creates a review for a movie by the user
func (s *MsServer) CreateReview(ctx context.Context, in *pb.Review) (*pb.Review, error) {
	newReview := schema.Review{
		Rating:      in.Rating,
		MovieID:     uint(in.MovieId),
		UserID:      uint(in.UserId),
		Description: in.Description,
	}
	err := s.Db.CreateReview(&newReview)
	if err != nil {
		log.Println(err)
	}
	return &pb.Review{
		ReviewId:    uint64(newReview.ID),
		Rating:      newReview.Rating,
		MovieId:     uint64(newReview.MovieID),
		UserId:      uint64(newReview.UserID),
		Description: newReview.Description,
	}, nil
}

// This RPC updates a review for a movie by the user
func (s *MsServer) UpdateReview(ctx context.Context, in *pb.Review) (*pb.Review, error) {
	review := schema.Review{
		Rating:      in.Rating,
		MovieID:     uint(in.MovieId),
		UserID:      uint(in.UserId),
		Description: in.Description,
	}
	err := s.Db.UpdateReview(in.ReviewId, review)
	if err != nil {
		log.Println(err)
		return &pb.Review{}, err
	}
	return &pb.Review{
		ReviewId:    uint64(in.ReviewId),
		UserId:      uint64(review.UserID),
		MovieId:     uint64(review.MovieID),
		Rating:      review.Rating,
		Description: review.Description,
	}, nil
}
func (s *MsServer) DeleteReview(ctx context.Context, in *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	err := s.Db.DeleteReview(in.GetUserId(), in.GetMovieId())
	if err != nil {
		return &pb.DeleteReviewResponse{
			UserId:  in.GetUserId(),
			MovieId: in.GetMovieId(),
			Status:  err.Error(),
		}, err
	}
	return &pb.DeleteReviewResponse{
		UserId:  in.GetUserId(),
		MovieId: in.GetMovieId(),
		Status:  "Review Deleted Successfully",
	}, nil
}

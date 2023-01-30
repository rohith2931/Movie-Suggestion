package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/utils"
	"log"
)

// This RPC adds movie into the watchlist
func (s *MsServer) AddMovieToWatchlist(ctx context.Context, in *pb.AddMovieToWatchlistRequest) (*pb.AddMovieToWatchlistResponse, error) {

	err := s.Db.AddMovieToWatchlist(in.UserId, in.MovieId)
	if err != nil {
		log.Println(err)
		return &pb.AddMovieToWatchlistResponse{}, err
	}
	return &pb.AddMovieToWatchlistResponse{
		MovieId: in.GetMovieId(),
		UserId:  in.GetUserId(),
		Status:  "Movie Added Successfully",
	}, nil
}

// This RPC gets all movies from the watchlist of a user
func (s *MsServer) GetAllWatchlistMovies(in *pb.GetAllWatchlistMoviesRequest, stream pb.MovieSuggestionService_GetAllWatchlistMoviesServer) error {

	Movies, err := s.Db.GetAllWatchlistMovies(in.UserId)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, movie := range Movies {
		Movie := &pb.Movie{
			MovieId:     uint64(movie.ID),
			Name:        movie.Name,
			Director:    movie.Director,
			Description: movie.Description,
			Rating:      movie.Rating,
			Language:    movie.Language,
			Category:    movie.Category,
			ReleaseDate: movie.ReleaseDate,
		}
		if err := stream.Send(Movie); err != nil {
			return err
		}
	}
	return nil
}

// This RPC deletes a movie from the watchlist
func (s *MsServer) DeleteMovieFromWatchlist(ctx context.Context, in *pb.DeleteMovieFromWatchlistRequest) (*pb.DeleteMovieFromWatchlistResponse, error) {

	err := s.Db.DeleteMovieFromWatchlist(in.UserId, in.MovieId)
	utils.CheckError(err)
	return &pb.DeleteMovieFromWatchlistResponse{
		UserId:  in.UserId,
		MovieId: in.MovieId,
		Status:  "Movie Deleted from watchlist successfully",
	}, nil
}

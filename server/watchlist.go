package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
)

// This RPC adds movie into the watchlist
func (s *MsServer) AddMovieToWatchlist(ctx context.Context, in *pb.AddMovieByUser) (*pb.Movie, error) {

	movie := schema.Movie{}
	movie = s.Db.AddMovieToWatchlist(in.UserId, in.MovieId)

	return &pb.Movie{
		Id:          uint64(movie.ID),
		Name:        movie.Name,
		Director:    movie.Director,
		Description: movie.Description,
		Rating:      uint64(movie.Rating),
		Language:    movie.Language,
		Category:    movie.Category,
		ReleaseDate: movie.ReleaseDate,
	}, nil
}

// This RPC gets all movies from the watchlist of a user
func (s *MsServer) GetAllWatchlistMovies(ctx context.Context, in *pb.UserId) (*pb.Movies, error) {

	Movies := s.Db.GetAllWatchlistMovies(in.Id)

	AllMovies := []*pb.Movie{}
	for _, movie := range Movies {
		AllMovies = append(AllMovies, &pb.Movie{
			Id:          uint64(movie.ID),
			Name:        movie.Name,
			Director:    movie.Director,
			Description: movie.Description,
			Rating:      uint64(movie.Rating),
			Language:    movie.Language,
			Category:    movie.Category,
			ReleaseDate: movie.ReleaseDate,
		})
	}
	return &pb.Movies{Movies: AllMovies}, nil
}

// This RPC deletes a movie from the watchlist
func (s *MsServer) DeleteMovieFromWatchlist(ctx context.Context, in *pb.DeleteMovieByUser) (*pb.Movie, error) {

	movie := schema.Movie{}
	movie = s.Db.DeleteMovieFromWatchlist(in.UserId, in.MovieId)
	Deletedmovie := &pb.Movie{
		Id:          uint64(movie.ID),
		Name:        movie.Name,
		Director:    movie.Director,
		Description: movie.Description,
		Rating:      uint64(movie.Rating),
		Language:    movie.Language,
		Category:    movie.Category,
		ReleaseDate: movie.ReleaseDate,
	}
	return Deletedmovie, nil
}

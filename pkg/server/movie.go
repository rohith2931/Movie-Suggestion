package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	rating "example/movieSuggestion/pkg/externalAPI"
	"example/movieSuggestion/pkg/schema"
	"log"
)

// This RPC returns all movies that are present in the database
func (s *MsServer) GetAllMovies(in *pb.MovieRequest, stream pb.MovieSuggestionService_GetAllMoviesServer) error {
	movies, err := s.Db.GetAllMovies(in.GetCategory(), in.GetLanguage())
	if err != nil {
		log.Println(err)
		return err
	}
	for _, movie := range movies {
		movie := &pb.Movie{
			MovieId:     uint64(movie.ID),
			Name:        movie.Name,
			Director:    movie.Director,
			Description: movie.Description,
			Rating:      movie.Rating,
			Language:    movie.Language,
			Category:    movie.Category,
			ReleaseDate: movie.ReleaseDate,
		}
		if err := stream.Send(movie); err != nil {
			return err
		}
	}
	return nil
}

// This RPC adds movie into the database
func (s *MsServer) AddMovie(ctx context.Context, in *pb.Movie) (*pb.Movie, error) {

	newMovie := schema.Movie{
		Name:        in.GetName(),
		Director:    in.GetDirector(),
		Description: in.GetDescription(),
		Language:    in.GetLanguage(),
		Category:    in.GetCategory(),
		ReleaseDate: in.GetReleaseDate(),
	}
	if in.GetRating() == 0 {
		newMovie.Rating = rating.GetImdbRating(newMovie.Name)

	} else {
		newMovie.Rating = in.GetRating()
	}
	err := s.Db.AddMovie(&newMovie)
	if err != nil {
		log.Println(err)
		return &pb.Movie{}, err
	}
	return &pb.Movie{
		MovieId:     uint64(newMovie.ID),
		Name:        newMovie.Name,
		Director:    newMovie.Director,
		Description: newMovie.Description,
		Rating:      newMovie.Rating,
		Language:    newMovie.Language,
		Category:    newMovie.Category,
		ReleaseDate: newMovie.ReleaseDate,
	}, nil
}

// This RPC Delete a movie from the database
func (s *MsServer) DeleteMovie(ctx context.Context, in *pb.Movie) (*pb.DeleteMovieResponse, error) {
	err := s.Db.DeleteMovie(in.MovieId)
	if err != nil {
		log.Println(err)
		return &pb.DeleteMovieResponse{}, err
	}

	return &pb.DeleteMovieResponse{
		MovieId: in.GetMovieId(),
		Status:  "Movie Deleted Successfully",
	}, nil
}

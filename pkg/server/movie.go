package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/pkg/rating"
	"example/movieSuggestion/pkg/schema"
	"example/movieSuggestion/utils"
)

// This RPC returns all movies that are present in the database
func (s *MsServer) GetAllMovies(in *pb.EmptyMovie, stream pb.MsDatabase_GetAllMoviesServer) error {
	Movies, err := s.Db.GetAllMovies()
	utils.CheckError(err)
	for _, movie := range Movies {
		Movie := &pb.Movie{
			Id:          uint64(movie.ID),
			Name:        movie.Name,
			Director:    movie.Director,
			Description: movie.Description,
			Rating:      uint64(movie.Rating),
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

// This RPC returns all movies based on category/genere
func (s *MsServer) GetMovieByCategory(ctx context.Context, in *pb.MovieCategory) (*pb.Movies, error) {
	AllMovies := []*pb.Movie{}
	Movies, err := s.Db.GetMovieByCategory(in.Category)
	utils.CheckError(err)
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

// This RPC adds movie into the database
func (s *MsServer) AddMovie(ctx context.Context, in *pb.NewMovie) (*pb.Movie, error) {

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
		newMovie.Rating = int(in.GetRating())
	}
	err := s.Db.AddMovie(&newMovie)
	utils.CheckError(err)
	return &pb.Movie{
		Name:        newMovie.Name,
		Director:    newMovie.Director,
		Description: newMovie.Description,
		Rating:      uint64(newMovie.Rating),
		Language:    newMovie.Language,
		Category:    newMovie.Category,
		ReleaseDate: newMovie.ReleaseDate,
		Id:          uint64(newMovie.ID)}, nil
}

// This RPC Delete a movie from the database
func (s *MsServer) DeleteMovie(ctx context.Context, in *pb.Movie) (*pb.Movie, error) {
	movie, err := s.Db.DeleteMovie(in.Id)
	utils.CheckError(err)

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

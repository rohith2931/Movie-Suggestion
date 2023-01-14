package server

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
	"log"
)

func (s *MsServer) AddMovieToWatchlist(ctx context.Context, in *pb.AddMovieByUser) (*pb.Movie, error) {
	log.Println("Add to Movie by user is called")

	watchlist := schema.Watchlist{}
	s.Db.Model(&schema.Watchlist{}).Where("user_id=?", in.UserId).Find(&watchlist)

	watchlistMoviesObj := schema.WatchlistMovies{
		WatchlistID: watchlist.ID,
		MovieID:     uint(in.MovieId),
	}

	s.Db.Create(&watchlistMoviesObj)

	// watchlistMovies := schema.WatchlistMovies{}
	// s.Db.Model(&schema.WatchlistMovies{}).Where("id=?", watchlist.ID).Find(&watchlistMovies)

	movie := schema.Movie{}
	// s.Db.Find(&movie, in.MovieId)

	// fmt.Println(watchlist.ID, movie.ID)
	// // fmt.Println(user.Watchlist)
	// // fmt.Println(user)
	// fmt.Printf("%v\n", watchlist)
	// // watchlist.Movies = append(watchlist.Movies, movie)
	s.Db.Find(&movie, in.MovieId)
	// // watchlist.Count = len(watchlist.Movies)
	// s.Db.Save(&watchlist)
	// fmt.Println()
	// fmt.Printf("%v\n", watchlist)

	// fmt.Println(user.Watchlist)
	// fmt.Println(user)

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

func (s *MsServer) GetAllWatchlistMovies(ctx context.Context, in *pb.UserId) (*pb.Movies, error) {
	log.Println("Get all movies from watchlist is called")

	watchlist := schema.Watchlist{}
	s.Db.Model(&schema.Watchlist{}).Where("user_id=?", in.Id).Find(&watchlist)

	Movies := []schema.Movie{}
	UserMovies := []schema.WatchlistMovies{}
	// row, _ := s.Db.DB().QueryContext(ctx, "SELECT movie_id FROM watchlist_movies WHERE watchlist_id=?", watchlist.ID)
	s.Db.Model(&schema.WatchlistMovies{}).Where("watchlist_id=?", watchlist.ID).Find(&UserMovies)
	// fmt.Printf("Length is %v", len(UserMovies))
	for _, UserMovie := range UserMovies {
		movie := schema.Movie{}
		s.Db.Model(&schema.Movie{}).Where("id=?", UserMovie.MovieID).Find(&movie)
		Movies = append(Movies, movie)
	}

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

func (s *MsServer) DeleteMovieFromWatchlist(ctx context.Context, in *pb.DeleteMovieByUser) (*pb.Movie, error) {
	log.Println("Delete Movie from watchlist by user is called")

	watchlist := schema.Watchlist{}
	s.Db.Model(&schema.Watchlist{}).Where("user_id=?", in.UserId).Find(&watchlist)
	watchlistMoviesObj := schema.WatchlistMovies{}
	s.Db.Model(&watchlistMoviesObj).Where("watchlist_id=? and movie_id=?", watchlist.ID, in.MovieId).Find(&watchlistMoviesObj)
	movie := schema.Movie{}
	s.Db.Find(&movie, watchlistMoviesObj.MovieID)
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
	s.Db.Model(&watchlistMoviesObj).Where("watchlist_id=? and movie_id=?", watchlist.ID, in.MovieId).Delete(&watchlistMoviesObj)
	return Deletedmovie, nil
}

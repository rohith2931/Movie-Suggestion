package database

import (
	"example/movieSuggestion/schema"

	"github.com/jinzhu/gorm"
)

type Database interface {
	CreateUser(*schema.User)
	GetAllMovies() []schema.Movie
	GetMovieByCategory(string) []schema.Movie
	AddMovie(*schema.Movie)
	DeleteMovie(uint64) schema.Movie
	CreateLike(uint64, uint64) string
	DeleteLike(uint64, uint64)
	UpdateReview(uint64, schema.Review)
	AddMovieToWatchlist(uint64, uint64)
	GetAllWatchlistMovies(uint64)
	DeleteMovieFromWatchlist()
}

type DBclient struct {
	Db *gorm.DB
}

func (db DBclient) CreateUser(newUser *schema.User) {
	db.Db.Create(newUser)
}

func (db DBclient) GetAllMovies() []schema.Movie {
	Movies := []schema.Movie{}
	db.Db.Find(&Movies)
	return Movies
}

func (db DBclient) GetMovieByCategory(Category string) []schema.Movie {
	Movies := []schema.Movie{}
	db.Db.Where("Category=?", Category).Find(&Movies)
	return Movies
}

func (db DBclient) AddMovie(newMovie *schema.Movie) {
	db.Db.Create(newMovie)
}

func (db DBclient) DeleteMovie(id uint64) schema.Movie {
	movie := schema.Movie{}
	db.Db.Find(&movie, id)
	db.Db.Unscoped().Delete(&schema.Movie{}, id)
	return movie
}

func (db DBclient) CreateLike(userId uint64, movieId uint64) string {
	likeObj := schema.Like{}
	db.Db.Where("user_id=? and movie_id=?", userId, movieId).Find(&likeObj)
	if likeObj.ID != 0 {
		return "User already liked the movie..."
	}
	db.Db.Debug().Create(&schema.Like{
		UserID:  uint(userId),
		MovieID: uint(movieId),
	})
	return "Like created Successfully.."
}

func (db DBclient) DeleteLike(userId uint64, movieId uint64) {
	db.Db.Unscoped().Model(&schema.Like{}).Where("user_id=? and movie_id=?", userId, movieId).Delete(&schema.Like{})
}

func (db DBclient) CreateReview(newReview *schema.Review) {
	db.Db.Create(newReview)
}

func (db DBclient) UpdateReview(id uint64, review schema.Review) {
	db.Db.Model(&schema.Review{}).Where("id=?", id).Updates(review)
}

func (db DBclient) AddMovieToWatchlist(userId uint64, movieId uint64) schema.Movie {
	watchlist := schema.Watchlist{}
	db.Db.Model(&schema.Watchlist{}).Where("user_id=?", userId).Find(&watchlist)

	watchlistMoviesObj := schema.WatchlistMovies{
		WatchlistID: watchlist.ID,
		MovieID:     uint(movieId),
	}

	db.Db.Create(&watchlistMoviesObj)
	movie := schema.Movie{}
	db.Db.Find(&movie, movieId)
	return movie
}

func (db DBclient) GetAllWatchlistMovies(userId uint64) []schema.Movie {
	Movies := []schema.Movie{}

	watchlist := schema.Watchlist{}
	db.Db.Model(&schema.Watchlist{}).Where("user_id=?", userId).Find(&watchlist)

	UserMovies := []schema.WatchlistMovies{}
	db.Db.Model(&schema.WatchlistMovies{}).Where("watchlist_id=?", watchlist.ID).Find(&UserMovies)
	for _, UserMovie := range UserMovies {
		movie := schema.Movie{}
		db.Db.Model(&schema.Movie{}).Where("id=?", UserMovie.MovieID).Find(&movie)
		Movies = append(Movies, movie)
	}
	return Movies
}

func (db DBclient) DeleteMovieFromWatchlist(userId uint64, movieId uint64) schema.Movie {
	watchlist := schema.Watchlist{}
	db.Db.Model(&schema.Watchlist{}).Where("user_id=?", userId).Find(&watchlist)
	watchlistMoviesObj := schema.WatchlistMovies{}
	db.Db.Model(&watchlistMoviesObj).Where("watchlist_id=? and movie_id=?", watchlist.ID, movieId).Find(&watchlistMoviesObj)
	movie := schema.Movie{}
	db.Db.Find(&movie, watchlistMoviesObj.MovieID)
	db.Db.Unscoped().Model(&watchlistMoviesObj).Where("watchlist_id=? and movie_id=?", watchlist.ID, movieId).Delete(&watchlistMoviesObj)
	return movie
}

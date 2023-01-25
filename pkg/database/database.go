package database

import (
	"example/movieSuggestion/pkg/schema"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Database interface {
	CreateUser(*schema.User) error
	GetAllMovies() ([]schema.Movie, error)
	GetMovieByCategory(string) ([]schema.Movie, error)
	AddMovie(*schema.Movie) error
	DeleteMovie(uint64) (schema.Movie, error)
	CreateLike(uint64, uint64) (string, error)
	DeleteLike(uint64, uint64) (string, error)
	CreateReview(*schema.Review) error
	UpdateReview(uint64, schema.Review) error
	AddMovieToWatchlist(uint64, uint64) (schema.Movie, error)
	GetAllWatchlistMovies(uint64) ([]schema.Movie, error)
	DeleteMovieFromWatchlist(uint64, uint64) (schema.Movie, error)
}

type DBclient struct {
	Db *gorm.DB
}

func (db DBclient) CreateUser(newUser *schema.User) error {
	err := db.Db.Create(newUser).Error
	return err
}

func (db DBclient) GetAllMovies() ([]schema.Movie, error) {
	Movies := []schema.Movie{}
	err := db.Db.Find(&Movies).Error
	return Movies, err
}

func (db DBclient) GetMovieByCategory(Category string) ([]schema.Movie, error) {
	Movies := []schema.Movie{}
	err := db.Db.Where("Category=?", Category).Find(&Movies).Error
	return Movies, err
}

func (db DBclient) AddMovie(newMovie *schema.Movie) error {
	err := db.Db.Create(newMovie).Error
	return err
}

func (db DBclient) DeleteMovie(id uint64) (schema.Movie, error) {
	movie := schema.Movie{}
	err := db.Db.Find(&movie, id).Error
	if err != nil {
		return movie, err
	}
	err = db.Db.Unscoped().Delete(&schema.Movie{}, id).Error
	return movie, err
}

func (db DBclient) CreateLike(userId uint64, movieId uint64) (string, error) {
	likeObj := schema.Like{}
	db.Db.Where("user_id=? and movie_id=?", userId, movieId).Find(&likeObj)
	if likeObj.ID != 0 {
		return "User already liked the movie...", nil
	}
	err := db.Db.Create(&schema.Like{
		UserID:  uint(userId),
		MovieID: uint(movieId),
	}).Error
	if err != nil {
		return err.Error(), err
	}
	return "Like created Successfully..", nil
}

func (db DBclient) DeleteLike(userId uint64, movieId uint64) (string, error) {
	err := db.Db.Unscoped().Model(&schema.Like{}).Where("user_id=? and movie_id=?", userId, movieId).Delete(&schema.Like{}).Error
	if err != nil {
		return err.Error(), err
	}
	return "Like Deleted Successfully..", nil
}

func (db DBclient) CreateReview(newReview *schema.Review) error {
	err := db.Db.Create(newReview).Error
	return err
}

func (db DBclient) UpdateReview(id uint64, review schema.Review) error {
	err := db.Db.Model(&schema.Review{}).Where("id=?", id).Updates(review).Error
	return err
}

func (db DBclient) AddMovieToWatchlist(userId uint64, movieId uint64) (schema.Movie, error) {
	movie := schema.Movie{}
	err := db.Db.Find(&movie, movieId).Error
	if err != nil {
		return movie, err
	}

	Watchlist := schema.Watchlist{}
	err = db.Db.Preload("WatchlistMovies").Model(&schema.Watchlist{}).Where("user_id=?", userId).Find(&Watchlist).Error
	Watchlist.WatchlistMovies = append(Watchlist.WatchlistMovies, schema.WatchlistMovies{MovieID: uint(movieId)})
	db.Db.Save(&Watchlist)
	return movie, err
}

func (db DBclient) GetAllWatchlistMovies(userId uint64) ([]schema.Movie, error) {

	Movies := []schema.Movie{}
	Watchlist := schema.Watchlist{}
	err := db.Db.Preload("WatchlistMovies").Model(&schema.Watchlist{}).Where("user_id=?", userId).Find(&Watchlist).Error
	fmt.Printf("%v", Watchlist.WatchlistMovies)
	for _, UserMovie := range Watchlist.WatchlistMovies {
		movie := schema.Movie{}
		db.Db.Model(&schema.Movie{}).Where("id=?", UserMovie.MovieID).Find(&movie)
		Movies = append(Movies, movie)
	}
	return Movies, err
}

func (db DBclient) DeleteMovieFromWatchlist(userId uint64, movieId uint64) (schema.Movie, error) {

	movie := schema.Movie{}
	err := db.Db.Find(&movie, movieId).Error
	if err != nil {
		return movie, err
	}
	watchlist := schema.Watchlist{}
	db.Db.Model(&schema.Watchlist{}).Where("user_id=?", userId).Find(&watchlist)
	db.Db.Unscoped().Model(&schema.WatchlistMovies{}).Where("watchlist_id=? and movie_id=?", watchlist.ID, movieId).Delete(&schema.WatchlistMovies{})
	return movie, err
}

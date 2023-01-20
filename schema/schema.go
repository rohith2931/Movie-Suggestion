package schema

import (
	"example/movieSuggestion/utils"

	"github.com/jinzhu/gorm"
)

func StartDB() {
	db, err := gorm.Open("postgres", utils.GoDotEnvVariable("DB_URL"))

	utils.CheckError(err)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Watchlist{})
	db.AutoMigrate(&WatchlistMovies{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Like{})

	db.Model(&Review{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Review{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE")
	db.Model(&Like{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Like{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE")
	db.Model(&Watchlist{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&WatchlistMovies{}).AddForeignKey("watchlist_id", "watchlists(id)", "CASCADE", "CASCADE")
	db.Model(&WatchlistMovies{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE")

	InitialiseDB(db)
}

type User struct {
	gorm.Model
	UserName    string
	Password    string
	Email       string
	PhoneNumber string
	Address     string
	Role        string
	Watchlist   *Watchlist
	Review      []Review
}
type Watchlist struct {
	gorm.Model
	UserID uint
}

type WatchlistMovies struct {
	gorm.Model
	WatchlistID uint
	MovieID     uint
}
type Review struct {
	gorm.Model
	Rating      int
	MovieID     uint
	UserID      uint
	Description string
}
type Movie struct {
	gorm.Model
	Name        string
	Director    string
	Description string
	Rating      int
	Language    string
	Category    string
	ReleaseDate string
}

type Like struct {
	gorm.Model
	UserID  uint
	MovieID uint
}

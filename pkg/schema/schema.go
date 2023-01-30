package schema

import (
	"github.com/jinzhu/gorm"
)

func StartDB(db *gorm.DB) {

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Watchlist{})
	db.AutoMigrate(&WatchlistMovies{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Like{})

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
	Reviews     []Review
	Likes       []Like
}
type Watchlist struct {
	gorm.Model
	UserID          uint
	WatchlistMovies []WatchlistMovies
}

type WatchlistMovies struct {
	gorm.Model
	WatchlistID uint
	MovieID     uint
}
type Review struct {
	gorm.Model
	Rating      float32
	MovieID     uint
	UserID      uint
	Description string
}
type Movie struct {
	gorm.Model
	Name            string
	Director        string
	Description     string
	Rating          float32
	Language        string
	Category        string
	ReleaseDate     string
	Likes           []Like
	Reviews         []Review
	WatchlistMovies []WatchlistMovies
}

type Like struct {
	gorm.Model
	UserID  uint
	MovieID uint
}

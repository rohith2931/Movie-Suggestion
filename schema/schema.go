package schema

import "github.com/jinzhu/gorm"

func StartDB() {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")

	CheckError(err)
	defer db.Close()

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

	db.Save(&Movie{
		Name:        "KGF",
		Director:    "Prashanth Neel",
		Description: "...",
		Rating:      9,
		Language:    "Telugu",
		Category:    "Action",
		ReleaseDate: "21-12-2018",
	})
	db.Save(&Movie{
		Name:        "F2",
		Director:    "Anil Ravipudi",
		Description: "Fun and frustation",
		Rating:      7,
		Language:    "Telugu",
		Category:    "Drama",
		ReleaseDate: "12-01-2019",
	})
	db.Debug().Save(&User{
		UserName:    "ram",
		Password:    "pass",
		Email:       "ram@gmail.com",
		PhoneNumber: "8932212342",
		Address:     "Hyderabad",
		Watchlist:   &Watchlist{},
	})
	// db.Save(&Watchlist{
	// 	UserID: 1,
	// })
	db.Save(&Review{
		Rating:      8,
		MovieID:     1,
		UserID:      1,
		Description: "Good",
	})
}

type User struct {
	gorm.Model
	UserName    string
	Password    string
	Email       string
	PhoneNumber string
	Address     string
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

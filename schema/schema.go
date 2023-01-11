package schema

import "github.com/jinzhu/gorm"

func StartDB() {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")

	CheckError(err)
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Watchlist{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Movie{})

	db.Model(&Review{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Review{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE")
	db.Model(&Watchlist{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Save(&Movie{
		Name:        "KGF",
		Director:    "Prashanth Neel",
		Description: "...",
		Rating:      9,
		Language:    "Telugu",
		Category:    "Action",
		ReleaseDate: "21-12-2018",
	})
	db.Debug().Save(&User{
		UserName:    "ram",
		Password:    "pass",
		Email:       "ram@gmail.com",
		PhoneNumber: "8932212342",
		Address:     "Hyderabad",
		Watchlist: Watchlist{
			Count: 0,
		},
	})
	// db.Save(&WatchList{
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
	Watchlist   Watchlist
	Review      []Review
}
type Watchlist struct {
	gorm.Model
	UserID uint
	Count  int
	Movies []Movie
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

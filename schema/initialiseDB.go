package schema

import "github.com/jinzhu/gorm"

func InitialiseDB(db *gorm.DB) {
	defer db.Close()
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

	//Creating admin
	db.Save(&User{
		UserName:    "rohith",
		Password:    "pass",
		Email:       "rohith@gmail.com",
		PhoneNumber: "8932212342",
		Address:     "Hyderabad",
		Role:        "admin",
		Watchlist:   &Watchlist{},
	})

	db.Save(&Review{
		Rating:      8,
		MovieID:     1,
		UserID:      1,
		Description: "Good",
	})
}

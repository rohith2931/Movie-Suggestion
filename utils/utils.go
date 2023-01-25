package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// This function checks for error
func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

// This function returns the value for the key from the .env file
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(filepath.Join("/home/rohith/Pet_Project/Movie Suggestion", ".env"))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

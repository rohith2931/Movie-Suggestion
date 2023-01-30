package rating

import (
	"encoding/json"
	"example/movieSuggestion/utils"
	"fmt"
	"io"
	"net/http"
)

// This function returns the Imdb rating of a movie
func GetImdbRating(movieName string) float32 {

	url := utils.GoDotEnvVariable("X-RapidAPI-URL")
	search := fmt.Sprintf("/?s=%s&l=1", movieName)
	req, _ := http.NewRequest("GET", url+search, nil)

	req.Header.Add("X-RapidAPI-Key", utils.GoDotEnvVariable("X-RapidAPI-Key"))
	req.Header.Add("X-RapidAPI-Host", utils.GoDotEnvVariable("X-RapidAPI-Host"))

	res, err := http.DefaultClient.Do(req)
	utils.CheckError(err)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var responseObj Response
	json.Unmarshal(body, &responseObj)

	rating := float32(responseObj.MovieRatings[0].Rating) / float32(10)
	return rating
}

type MovieRating struct {
	Rating int `json:"score_average"`
}
type Response struct {
	MovieRatings []MovieRating `json:"search"`
}

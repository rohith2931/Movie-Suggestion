package rating

import (
	"encoding/json"
	"example/movieSuggestion/utils"
	"fmt"
	"io"
	"net/http"
)

// This function returns the Imdb rating of a movie
func GetImdbRating(movieName string) int {

	url := fmt.Sprintf("https://mdblist.p.rapidapi.com/?s=%s&l=1", movieName)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "2636fc2a6cmshe12aded30b561dcp132388jsn7f2610a57bd0")
	req.Header.Add("X-RapidAPI-Host", "mdblist.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	utils.CheckError(err)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var responseObj Response
	json.Unmarshal(body, &responseObj)

	rating := responseObj.MovieRatings[0].Rating / 10
	return rating
}

type MovieRating struct {
	Rating int `json:"score_average"`
}
type Response struct {
	MovieRatings []MovieRating `json:"search"`
}

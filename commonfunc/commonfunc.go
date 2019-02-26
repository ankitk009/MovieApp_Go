package commonfunc

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//Movie struct consist of skeleton of required data
type Movie struct {
	Title     string       `json:"Title"`
	Movietype string       `json:"Type"`
	Rating    Ratingstruct `json:"Rating"`
	Year      string       `json:"Year"`
	Genre     string       `json:"Genre"`
	Plot      string       `json:"Description"`
	Runtime   string       `json:"Runtime"`
	Actors    string       `json:"Actors"`
	Country   string       `json:"Country"`
}

//Ratingstruct struct which incldes various rating agencies
type Ratingstruct struct {
	Imdb           string `json:"Imdb"`
	Rottentamotoes string `json:"RottenTomatoes"`
	Metacritic     string `json:"Metacritic"`
}

//MovieList struct includes the list of all the movies
type MovieList struct {
	Title string `json:"Title"`
	Type  string `json:"Type"`
	Year  string `json:"Year"`
}

//ProcessSearchResult function processes the result and returns a Movie object
func ProcessSearchResult(results map[string]interface{}) Movie {
	var movie Movie
	var ok bool
	movie.Title, ok = results["Title"].(string)
	if !ok {
		return movie
	}
	movie.Movietype = results["Type"].(string)
	movie.Year = results["Year"].(string)

	movie.Genre = results["Genre"].(string)
	movie.Plot = results["Plot"].(string)
	movie.Runtime = results["Runtime"].(string)
	movie.Actors = results["Actors"].(string)
	movie.Country = results["Country"].(string)

	ratings := results["Ratings"].([]interface{})

	if len(ratings) == 3 {
		movie.Rating.Imdb = results["Ratings"].([]interface{})[0].(map[string]interface{})["Value"].(string)
		movie.Rating.Rottentamotoes = results["Ratings"].([]interface{})[1].(map[string]interface{})["Value"].(string)
		movie.Rating.Metacritic = results["Ratings"].([]interface{})[2].(map[string]interface{})["Value"].(string)
	} else if len(ratings) == 2 {
		movie.Rating.Imdb = results["Ratings"].([]interface{})[0].(map[string]interface{})["Value"].(string)
		movie.Rating.Rottentamotoes = results["Ratings"].([]interface{})[1].(map[string]interface{})["Value"].(string)
		movie.Rating.Metacritic = "N/A"
	} else if len(ratings) == 1 {
		movie.Rating.Imdb = results["Ratings"].([]interface{})[0].(map[string]interface{})["Value"].(string)
		movie.Rating.Rottentamotoes = "N/A"
		movie.Rating.Metacritic = "N/A"
	}

	// fmt.Println(movie)
	return movie
}

//SendAndReceiveRequest function sends the http request and receives the request body which it returns
func SendAndReceiveRequest(baseurl *url.URL) []byte {
	movieclient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, baseurl.String(), nil)
	Validate(err)

	req.Header.Set("User-Agent", "Smalltutorial")

	res, getErr := movieclient.Do(req)

	Validate(getErr)

	body, readErr := ioutil.ReadAll(res.Body)
	Validate(readErr)

	return body
}

//GenerateBaseURL creates an API Base URL and returns the encoded URL
func GenerateBaseURL() *url.URL {
	apikey := "579571ec"
	//Creating base URL
	baseurl, _ := url.Parse("http://www.omdbapi.com/")
	v := baseurl.Query()
	v.Set("apikey", apikey)
	baseurl.RawQuery = v.Encode()
	return baseurl
}

//Validate function logs the error
func Validate(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//ProcessListResults takes the result and maps into MovieList object
func ProcessListResults(result map[string]interface{}) []MovieList {

	var movielist []MovieList
	if result["Response"].(string) == "False" {
		return movielist
	}
	length, err := strconv.Atoi(result["totalResults"].(string))
	if length > 10 && err != nil {
		length = 10
	}
	movielist = make([]MovieList, length)

	//fmt.Println(result)
	movieArray := result["Search"].([]interface{})
	for key, value := range movieArray {
		movielist[key].Title = value.(map[string]interface{})["Title"].(string)
		movielist[key].Type = value.(map[string]interface{})["Type"].(string)
		movielist[key].Year = value.(map[string]interface{})["Year"].(string)
	}
	return movielist
}

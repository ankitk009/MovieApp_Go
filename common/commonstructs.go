package common

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

package structure

type PopularFilmsData struct {
	Title        string
	Vote_Average int
	Vote_Count   int
	Overview     string
	Release_date string
	Poster_path  string
}

type ApiData struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type ForFavs struct {
	user  string `json:"user"`
	Title string `json:"movie"`
}

type Films struct {
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	PosterPath  string  `json:"poster_path"`
	VoteAverage float64 `json:"vote_average"`
}

type All struct {
	Results []PopularFilmsData `json:"results"`
}

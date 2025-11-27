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

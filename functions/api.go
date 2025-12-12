package functions

import (
	"GroupieTracker/structure"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

func GetPopularFilms() ([]structure.PopularFilmsData, error) {

	//Structure qui prend les données JSON
	type ApiData struct {
		Results []struct {
			Title        string  `json:"title"`
			Vote_Average float32 `json:"vote_Average"`
			Vote_Count   int     `json:"vote_Count"`
			Overview     string  `json:"overview"`
			Release_date string  `json:"release_date"`
			Poster_path  string  `json:"Poster_path"`
		}
	}

	films := []structure.PopularFilmsData{}

	// URL de L'API
	urlApi := "https://api.themoviedb.org/3/movie/popular?language=fr-FR"

	// Initialisation du client HTTP qui va émettre/demander les requêtes avec un temps d'arrêt après 2 secondes
	httpClient := http.Client{
		Timeout: time.Second * 10, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête, retourne rien si pas d'erreur, sinon retourne une erreur
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return nil, errReq
	}

	// Permet d'ajouter des infos au Header de la requête, ici le type de token et le token
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJmZjU2MTA5NDEwNTJjOTFjMDUxN2Q0M2JkZmQ1MzY1ZSIsIm5iZiI6MTc2MDMwMTAwOC4wNzMsInN1YiI6IjY4ZWMwZmQwMzNkMWFlYTViOWE3MjZlZiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.ROh7GMpu3OqR6J1yM1sabgrt2ziwPc2e-38wnbjqMww")

	// Execution de la requête HTTP vers L'API, retourne rien si pas d'erreur, sinon retourne une erreur
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP, retourne rien si pas d'erreur, sinon retourne une erreur
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		return nil, errBody
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	//Boucle qui remplit popularFilms avec les données récupérées dans la structure ApiData qu'on a décodé depuis le JSON
	for i := 0; i <= len(decodeData.Results)-1; i++ {
		var popularFilms structure.PopularFilmsData
		popularFilms.Title = decodeData.Results[i].Title
		popularFilms.VoteAverage = decodeData.Results[i].Vote_Average
		popularFilms.VoteCount = decodeData.Results[i].Vote_Count
		popularFilms.Overview = decodeData.Results[i].Overview
		popularFilms.ReleaseDate = decodeData.Results[i].Release_date
		popularFilms.PosterPath = "https://image.tmdb.org/t/p/original" + decodeData.Results[i].Poster_path

		films = append(films, popularFilms)
	}

	return films, nil

}

func GetSearchFilm(query string) ([]structure.PopularFilmsData, error) {
	//Structure qui prend les données JSON
	type ApiData struct {
		Results []struct {
			Title        string  `json:"title"`
			Vote_Average float32 `json:"vote_Average"`
			Vote_Count   int     `json:"vote_Count"`
			Overview     string  `json:"overview"`
			Release_date string  `json:"release_date"`
			Poster_path  string  `json:"Poster_path"`
		}
	}

	films := []structure.PopularFilmsData{}

	// url.QueryEscape(query) premet d'encoder la string pour pas avoir d'erreur avec les espace, accents, etc
	// URL de L'API
	urlApi := "https://api.themoviedb.org/3/search/movie?query=" + url.QueryEscape(query) + "&include_adult=false&language=fr-FR"

	// Initialisation du client HTTP qui va émettre/demander les requêtes avec un temps d'arrêt après 2 secondes
	httpClient := http.Client{
		Timeout: time.Second * 5, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête, retourne rien si pas d'erreur, sinon retourne une erreur
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return nil, errReq
	}

	// Permet d'ajouter des infos au Header de la requête, ici le type de token et le token
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJmZjU2MTA5NDEwNTJjOTFjMDUxN2Q0M2JkZmQ1MzY1ZSIsIm5iZiI6MTc2MDMwMTAwOC4wNzMsInN1YiI6IjY4ZWMwZmQwMzNkMWFlYTViOWE3MjZlZiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.ROh7GMpu3OqR6J1yM1sabgrt2ziwPc2e-38wnbjqMww")

	// Execution de la requête HTTP vers L'API, retourne rien si pas d'erreur, sinon retourne une erreur
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP, retourne rien si pas d'erreur, sinon retourne une erreur
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		return nil, errBody
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	//Boucle qui remplit popularFilms avec les données récupérées dans la structure ApiData qu'on a décodé depuis le JSON
	//Le if permet d'exclure les films qui n'ont pas de poster et pas de résumé
	for i := 0; i < len(decodeData.Results); i++ {
		var popularFilms structure.PopularFilmsData
		if decodeData.Results[i].Poster_path != "" && decodeData.Results[i].Overview != "" {
			popularFilms.Title = decodeData.Results[i].Title
			popularFilms.VoteAverage = decodeData.Results[i].Vote_Average
			popularFilms.VoteCount = decodeData.Results[i].Vote_Count
			popularFilms.Overview = decodeData.Results[i].Overview
			popularFilms.ReleaseDate = decodeData.Results[i].Release_date
			popularFilms.PosterPath = "https://image.tmdb.org/t/p/original" + decodeData.Results[i].Poster_path
			films = append(films, popularFilms)

		}
	}
	return films, nil
}

/*func SetFavouriteFilm(filename string, fav structure.Favourite) error {
	err := SaveFavouriteFile(filename, fav)
	if err != nil {
		return err
	}
	return nil
}*/

package controller

import (
	"GroupieTracker/functions"
	"GroupieTracker/structure"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

// Initialise la structure PageData qui sera chargée dans chaque template HTML
type PageData struct {
	PopularFilms  []structure.PopularFilmsData
	PopularActors []structure.PopularActors
}

// Fonction qui gère le template de la page Accueil
func Index(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

// Fonction qui gère le template de la page À Propos
func About(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/about.html"))
	tmpl.Execute(w, data)
}

// Fonction qui gère le template de la page Categorie
func Categorie(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/categorie.html"))
	tmpl.Execute(w, data)
}

// Fonction qui gère le template de la page Collection
func Collection(w http.ResponseWriter, r *http.Request) {
	films, err := functions.GetPopularFilms()
	if err != nil {
		log.Println("Erreur récupération films:", err)
		return
	}
	data := PageData{
		PopularFilms: films,
	}
	tmpl := template.Must(template.ParseFiles("template/collection.html"))
	tmpl.Execute(w, data)
}

// Fonction qui gère le template de la page Ressources
func Ressources(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/ressources.html"))
	tmpl.Execute(w, data)
}

// Fonction qui gère le template de la page Favoris
func Favoris(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/favoris.html"))
	tmpl.Execute(w, data)
}

func Acteurs(w http.ResponseWriter, r *http.Request) {
	actors, err := functions.GetActors()
	if err != nil {
		log.Println("Erreur récupération films:", err)
		return
	}
	data := PageData{
		PopularActors: actors,
	}
	tmpl := template.Must(template.ParseFiles("template/Acteurs.html"))
	tmpl.Execute(w, data)
}

// Fonction qui gère le template de la page Login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // Si le formulaire est soumis en POST
		// Récupération des données du formulaire
		name := r.FormValue("name") // Récupère le champ "name"
		msg := r.FormValue("msg")   // Récupère le champ "msg"
		data := map[string]string{
			"Title":   "Contact",
			"Message": "Merci " + name + " pour ton message : " + msg, // Message personnalisé après soumission
		}
		renderTemplate(w, "login.html", data)
		return // On termine ici pour ne pas exécuter la partie GET

	}
	// Si ce n'est pas un POST, on affiche simplement le formulaire
	data := map[string]string{
		"Title":   "",
		"Message": "Se connecter",
	}
	renderTemplate(w, "login.html", data)
}

// Fonction qui gère le template de la page Signup
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // Si le formulaire est soumis en POST
		// Récupération des données du formulaire
		name := r.FormValue("name") // Récupère le champ "name"
		msg := r.FormValue("msg")   // Récupère le champ "msg"
		data := map[string]string{
			"Title":   "Contact",
			"Message": "Merci " + name + " pour ton message : " + msg, // Message personnalisé après soumission
		}
		renderTemplate(w, "signup.html", data)
		return // On termine ici pour ne pas exécuter la partie GET

	}
	// Si ce n'est pas un POST, on affiche simplement le formulaire
	data := map[string]string{
		"Title":   "",
		"Message": "S'inscrire",
	}
	renderTemplate(w, "signup.html", data)
}

// Fonction qui gère le template de la page Recherche
func Recherche(w http.ResponseWriter, r *http.Request) {

	var searchedfilms []structure.PopularFilmsData

	if r.Method == http.MethodPost {
		name := r.FormValue("search-query")

		films, err := functions.GetSearchFilm(name)
		searchedfilms = films
		if err != nil {
			log.Println("Erreur récupération films:", err)
			return

		}
	}

	data := PageData{
		PopularFilms: searchedfilms,
	}

	tmpl := template.Must(template.ParseFiles("template/recherche.html"))
	tmpl.Execute(w, data)

}

func AddFavoris(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // Vérifie que la méthode est POST
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	title := r.FormValue("titre") // Récupère le titre du film depuis le formulaire
	overview := r.FormValue("overview")
	release_date := r.FormValue("release_date")
	poster_path := r.FormValue("poster_path")
	voteAveragestr := r.FormValue("vote_average")
	voteAverage, err := strconv.ParseFloat(voteAveragestr, 64)
	if err != nil {
		log.Println("Erreur conversion vote_average:", err)
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	// Lit le fichier JSON des favoris existants

	data, err := os.ReadFile("data/favourite.json")
	if err != nil {
		log.Println("Erreur lecture fichier favoris:", err)
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	var favs []structure.Films
	if len(data) > 0 {
		if err := json.Unmarshal(data, &favs); err != nil {
			log.Println("Erreur décodage JSON favoris:", err)
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
	}

	found := false
	newFavs := []structure.Films{}
	for _, f := range favs {
		if f.Title == title {
			found = true
			continue
		}
		newFavs = append(newFavs, f)
	}

	if !found {
		newID := 1
		if len(favs) > 0 {
			newID = favs[len(favs)-1].Id + 1
		}

		newFav := structure.Films{
			Id:          newID,
			Title:       title,
			Overview:    overview,
			ReleaseDate: release_date,
			PosterPath:  poster_path,
			VoteAverage: voteAverage,
		}
		newFavs = append(newFavs, newFav)
		log.Println("Film ajouté aux favoris:", title)
	} else {
		log.Println("Film retiré des favoris:", title)
	}

	updatedData, err := json.MarshalIndent(newFavs, "", "  ")
	if err != nil {
		log.Println("Erreur écriture JSON favoris:", err)
	}

	if err := os.WriteFile("data/favourite.json", updatedData, 0644); err != nil {
		log.Println("Erreur écriture fichier favoris:", err)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

// Télécharge les films populaires depuis l'API TMDB
func GetMovies() ([]structure.PopularFilmsData, error) {
	apiKey := "ff5610941052c91c0517d43bdfd5365e"
	urlApi := "https://api.themoviedb.org/3/movie/popular?&language=fr-FR&api_key=" + apiKey

	resp, err := http.Get(urlApi)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data structure.All
	json.Unmarshal(body, &data)

	return data.Results, nil
}

func findFilm(movies []structure.PopularFilmsData, title string) *structure.PopularFilmsData {
	title = (title)

	for _, m := range movies {
		if (m.Title) == title {
			return &m
		}
	}

	return nil
}

func ShowFavs(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data/favourite.json")
	if err != nil {
		http.Error(w, "Erreur serveur", 500)
		return
	}

	var favs []structure.Films
	if len(data) > 0 {
		json.Unmarshal(data, &favs)
	}

	tmpl, err := template.ParseFiles("template/favoris.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}

	tmpl.Execute(w, favs)
}

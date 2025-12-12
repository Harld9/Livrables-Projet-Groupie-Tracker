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
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

// Initialise la structure PageData qui sera chargée dans chaque template HTML
type PageData struct {
	PopularFilms []structure.PopularFilmsData
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
	if r.Method == http.MethodPost {
		movie := r.FormValue("titre")
		log.Println("ajout du film:", movie, "aux favoris.")

		//Lire les favoris existants
		data, err := os.ReadFile("data/favourite.json")
		if err != nil {
			log.Println("Erreur lecture fichier favoris:", err)
			return
		}

		var favs []structure.ForFavs
		if len(data) > 0 {
			if err := json.Unmarshal(data, &favs); err != nil {
				log.Println("Erreur décodage JSON favoris:", err)
				return
			}
		}
		//Ajouter le nouveau favori
		favs = append(favs, structure.ForFavs{Title: movie})

		//Enregistrer les favoris mis à jour
		updatedData, err := json.MarshalIndent(favs, "", "  ")
		if err != nil {
			log.Println("Erreur écriture fichier favoris:", err)
			return
		}

		err = os.WriteFile("data/favourite.json", updatedData, 0644)
		if err != nil {
			log.Println("Erreur écriture fichier favoris:", err)
			return
		}
		log.Println("Film ajouté aux favoris avec succès.")
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

	data, _ := os.ReadFile("data/favourite.json")

	var favs []structure.ForFavs
	json.Unmarshal(data, &favs)

	popular, _ := GetMovies()

	selectedFav := []structure.PopularFilmsData{}

	for _, fav := range favs {
		movie := findFilm(popular, fav.Title)
		if movie != nil {
			selectedFav = append(selectedFav, *movie)
		}
	}

	tmpl := template.Must(template.ParseFiles("template/favoris.html"))
	tmpl.Execute(w, selectedFav)
}

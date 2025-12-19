package controller

import (
	"GroupieTracker/functions"
	"GroupieTracker/structure"
	"encoding/json"
	"html/template"
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
		// Récupérer les valeurs du formulaire
		title := r.FormValue("title")
		overview := r.FormValue("overview")
		releaseDate := r.FormValue("release_date")
		posterPath := r.FormValue("poster_path")
		voteAverageStr := r.FormValue("vote_average")

		// Convertir voteAverage en float64
		voteAverage, err := strconv.ParseFloat(voteAverageStr, 64)
		if err != nil {
			log.Println("Erreur conversion vote_average:", err)
			voteAverage = 0.0
		}

		// Lire les favoris existants
		data, err := os.ReadFile("data/favourite.json")
		if err != nil {
			log.Println("Erreur lecture fichier favoris:", err)
			return
		}

		var favs []structure.Films
		if len(data) > 0 {
			if err := json.Unmarshal(data, &favs); err != nil {
				log.Println("Erreur décodage JSON favoris:", err)
				return
			}
		}

		// Déterminer le nouvel ID
		newID := 1
		if len(favs) > 0 {
			newID = favs[len(favs)-1].Id + 1
		}

		// Ajouter le nouveau favori
		newFilm := structure.Films{
			Id:          newID,
			Title:       title,
			Overview:    overview,
			ReleaseDate: releaseDate,
			PosterPath:  posterPath,
			VoteAverage: voteAverage,
		}
		favs = append(favs, newFilm)

		// Enregistrer les favoris mis à jour
		updatedData, err := json.MarshalIndent(favs, "", "  ")
		if err != nil {
			log.Println("Erreur écriture JSON favoris:", err)
			return
		}

		err = os.WriteFile("data/favourite.json", updatedData, 0644)
		if err != nil {
			log.Println("Erreur écriture fichier favoris:", err)
			return
		}

		log.Println("Film ajouté aux favoris avec succès:", newFilm.Title)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
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

package controller

import (
	"GroupieTracker/functions"
	"GroupieTracker/structure"
	"html/template"
	"log"
	"net/http"
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

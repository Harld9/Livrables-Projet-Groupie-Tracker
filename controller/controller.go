package controller

import (
	"html/template"
	"net/http"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

type PageData struct {
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func About(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/about.html"))
	tmpl.Execute(w, data)
}

func Categorie(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/categorie.html"))
	tmpl.Execute(w, data)
}

func Collection(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/collection.html"))
	tmpl.Execute(w, data)
}

func Recherche(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/recherche.html"))
	tmpl.Execute(w, data)
}

func Ressources(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/ressources.html"))
	tmpl.Execute(w, data)
}

func Favoris(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/favoris.html"))
	tmpl.Execute(w, data)
}

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

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

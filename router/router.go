package router

import (
	"GroupieTracker/controller"
	"net/http"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'application
func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Routes de ton app
	mux.HandleFunc("/", controller.Index)
	mux.HandleFunc("/about", controller.About)
	mux.HandleFunc("/categorie", controller.Categorie)
	mux.HandleFunc("/collection", controller.Collection)
	mux.HandleFunc("/ressources", controller.Ressources)
	mux.HandleFunc("/favoris", controller.Favoris)
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/signup", controller.Signup)
	mux.HandleFunc("/recherche", controller.Recherche)

	// Ajout des fichiers statiques
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}

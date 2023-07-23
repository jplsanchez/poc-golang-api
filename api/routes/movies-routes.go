package routes

import (
	"github.com/gorilla/mux"
	"github.com/jplsanchez/poc-golang-api/api/controllers"
)

var RegisterMoviesRoutes = func(router *mux.Router, controller *controllers.MovieController) {
	router.HandleFunc("/movies", controller.GetMovies).Methods("GET")
	router.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controller.GetMovieById).Methods("GET")
	router.HandleFunc("/movies/{id}", controller.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controller.DeleteMovie).Methods("DELETE")
}

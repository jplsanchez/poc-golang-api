package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jplsanchez/poc-golang-api/api/controllers"
	"github.com/jplsanchez/poc-golang-api/api/routes"
	"github.com/jplsanchez/poc-golang-api/data/repositories"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterMoviesRoutes(router, MovieDiHandler())
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:5000", router))
}

func MovieDiHandler() *controllers.MovieController {
	movieRepository := repositories.NewMovieRepository()
	movieController := controllers.NewMovieController(movieRepository)
	return movieController
}

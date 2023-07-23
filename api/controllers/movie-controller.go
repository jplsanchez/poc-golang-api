package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	utils "github.com/jplsanchez/poc-golang-api/api/controllers/common"
	"github.com/jplsanchez/poc-golang-api/movies"
	"github.com/jplsanchez/poc-golang-api/movies/models"
)

type MovieController struct {
	repository movies.Repository
}

func NewMovieController(repository movies.Repository) *MovieController {
	c := MovieController{}
	c.repository = repository
	return &c
}

func (c MovieController) GetMovies(writer http.ResponseWriter, request *http.Request) {
	movies := c.repository.GetAllMovies()
	utils.WriteOkResponse[[]models.Movie](movies, writer)
}

func (c MovieController) GetMovieById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	movieId := params["id"]
	id, err := uuid.Parse(movieId)
	if err != nil {
		log.Println("Error while parsing")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	movie := c.repository.GetMovieById(id)
	if movie.IsEmptyOrNil() {
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	utils.WriteOkResponse[models.Movie](movie, writer)
}

func (c MovieController) CreateMovie(writer http.ResponseWriter, request *http.Request) {
	movie := &models.Movie{}
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.Id = uuid.New()
	movie = c.repository.CreateMovie(movie)
	utils.WriteOkResponse[models.Movie](movie, writer)
}

func (c MovieController) UpdateMovie(writer http.ResponseWriter, request *http.Request) {
	movie := &models.Movie{}
	_ = json.NewDecoder(request.Body).Decode(&movie)
	params := mux.Vars(request)
	movieId := params["id"]
	id, err := uuid.Parse(movieId)
	if err != nil {
		log.Println("Error while parsing")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	movie = c.repository.UpdateMovie(id, movie)
	if movie.IsEmptyOrNil() {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("ID was not found.")
		return
	}
	utils.WriteOkResponse[models.Movie](movie, writer)
}

func (c MovieController) DeleteMovie(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	movieId := params["id"]
	id, err := uuid.Parse(movieId)
	if err != nil {
		log.Println("Error while parsing")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	movie := c.repository.DeleteMovie(id)

	if movie.IsEmptyOrNil() {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("ID was not found.")
		return
	}
	utils.WriteOkResponse[models.Movie](movie, writer)
}

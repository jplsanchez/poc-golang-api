package movies

import (
	"github.com/google/uuid"
	"github.com/jplsanchez/poc-golang-api/movies/models"
)

type Repository interface {
	CreateMovie(movie *models.Movie) *models.Movie
	GetAllMovies() *[]models.Movie
	GetMovieById(id uuid.UUID) *models.Movie
	DeleteMovie(id uuid.UUID) *models.Movie
	UpdateMovie(id uuid.UUID, updateMovie *models.Movie) *models.Movie
}

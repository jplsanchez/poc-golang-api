package repositories

import (
	"github.com/google/uuid"
	"github.com/jplsanchez/poc-golang-api/data/config"
	"github.com/jplsanchez/poc-golang-api/movies/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	Database *gorm.DB
}

func NewMovieRepository() *MovieRepository {
	r := MovieRepository{}
	config.Connect()
	r.Database = config.GetDatabase()
	r.Database.AutoMigrate(&models.Movie{})
	return &r
}

func (r MovieRepository) CreateMovie(movie *models.Movie) *models.Movie {
	r.Database.Create(movie)
	return movie
}

func (r MovieRepository) GetAllMovies() *[]models.Movie {
	var movies []models.Movie
	r.Database.Find(&movies)
	return &movies
}

func (r MovieRepository) GetMovieById(id uuid.UUID) *models.Movie {
	var movie models.Movie
	r.Database.Where("Id=?", id).Find(&movie)
	return &movie
}

func (r MovieRepository) DeleteMovie(id uuid.UUID) *models.Movie {
	var movie models.Movie
	r.Database.Where("Id=?", id).Find(&movie).Delete(movie)
	return &movie
}

func (r MovieRepository) UpdateMovie(id uuid.UUID, updateMovie *models.Movie) *models.Movie {
	var movie models.Movie
	db := r.Database.Where("Id=?", id).Find(&movie)
	if movie.Id == uuid.Nil {
		return nil
	}
	if updateMovie.Director != "" {
		movie.Director = updateMovie.Director
	}
	if updateMovie.Title != "" {
		movie.Title = updateMovie.Title
	}
	db.Save(&movie)
	return &movie
}

package data

import "github.com/skufu/movies/models"

type MovieStorage interface {
	GetTopMovies() ([]models.Movie, error)
	GetRandomMovies() ([]models.Movie, error)
	//GetMovieByID(id int) (models.Movie, error)
	//SearchMoviesByName(name string) ([]models.Movie, error)
	//GetAllGenre() ([]models.Movie, error)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/skufu/movies/data"
	"github.com/skufu/movies/logger"
	"github.com/skufu/movies/models"
)

type MovieHandler struct {
	Storage data.MovieStorage // ← Use interface instead of concrete type
	Logger  *logger.Logger
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Logger.Error("Error encoding JSON", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	//calls get top movies from MovieRepository
	movies, err := h.Storage.GetTopMovies()
	if err != nil {
		h.Logger.Error("Error getting top movies", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     123,
			Title:       "The Dark Knight",
			Tagline:     "The Dark Knight",
			ReleaseYear: 2008,
			Genre: []models.Genre{
				{ID: 1, Name: "Action"},
			},
			Keywords: []string{},
			Casting:  []models.Actor{{ID: 1, FirstName: "Max", LastName: "Max"}},
		},
		{
			ID:          2,
			TMDB_ID:     124,
			Title:       "How to train your dragon",
			Tagline:     "How to train your dragon",
			ReleaseYear: 2010,
			Genre: []models.Genre{
				{ID: 2, Name: "Animation"},
			},
			Keywords: []string{},
			Casting:  []models.Actor{{ID: 2, FirstName: "John", LastName: "Doe"}},
		},
	}
	h.writeJSONResponse(w, movies)

}

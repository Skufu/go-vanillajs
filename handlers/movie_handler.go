package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/skufu/movies/models"
)

type MovieHandler struct {
	//todo
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
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
			Casting:  []models.Actor{{ID: 1, Name: "Max"}},
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
		},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		//log error

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (h *MovieHandler) handleStorageError(w http.ResponseWriter, err error, message string) bool {
	if err != nil {
		h.Logger.Error(message, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func (h *MovieHandler) parseID(w http.ResponseWriter, idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error("Invalid ID format", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
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
	movies, err := h.Storage.GetRandomMovies()
	if err != nil {
		h.Logger.Error("Error getting random movies", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	order := r.URL.Query().Get("order")
	genreStr := r.URL.Query().Get("genre")

	var genre *int
	if genreStr != "" {
		genreInt, ok := h.parseID(w, genreStr)
		if !ok {
			return
		}
		genre = &genreInt
	}

	var movies []models.Movie
	var err error
	if query != "" {
		movies, err = h.Storage.SearchMoviesByName(query, order, genre)
	}
	if h.handleStorageError(w, err, "Failed to get movies") {
		return
	}
	h.writeJSONResponse(w, movies)
	h.Logger.Info("Successfully served movies")
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/movies/"):]
	id, ok := h.parseID(w, idStr)
	if !ok {
		return
	}

	movie, err := h.Storage.GetMovieByID(id)
	if h.handleStorageError(w, err, "Failed to get movie by ID") {
		return
	}
	h.writeJSONResponse(w, movie)
	h.Logger.Info("Successfully served movie with ID: " + idStr)
}

func (h *MovieHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.Storage.GetAllGenres()
	if h.handleStorageError(w, err, "Failed to get genres") {
		return
	}
	h.writeJSONResponse(w, genres)
	h.Logger.Info("Successfully served genres")
}

func NewMovieHandler(storage data.MovieStorage, log *logger.Logger) *MovieHandler {
	return &MovieHandler{
		Storage: storage,
		Logger:  log,
	}
}

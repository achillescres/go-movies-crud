package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-movies-crud/handler/check"
	"go-movies-crud/models"
	"go-movies-crud/repository"
	"net/http"
)

type Handler struct {
	*repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{Repository: repository}
}

func (handler *Handler) GetMovies(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies, err := handler.Repository.Movies.GetMovies()
	check.Error(w, err)

	err = json.NewEncoder(w).Encode(movies)
	check.Error(w, err)
}

func (handler *Handler) GetMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	movie, err := handler.Repository.Movies.GetMovie(id)
	if check.Error(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(movie)
	check.Error(w, err)
}

func (handler *Handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ok, err := handler.Repository.Movies.DeleteMovie(id)

	if check.Existence(w, ok, id) {
		return
	}
	check.Error(w, err)
}

func (handler *Handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		check.RaiseInternalError(w, errors.New("request hasn't \"id\" param"))
		return
	}

	var newMovie models.Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if check.Error(w, err) {
		return
	}

	ok, err = handler.Repository.Movies.UpdateMovie(id, newMovie)
	if check.Existence(w, ok, id) {
		return
	}
	check.Error(w, err)
}

func (handler *Handler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie models.Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if check.Error(w, err) {
		return
	}

	err = handler.Repository.Movies.CreateMovie(newMovie)
	check.Error(w, err)
}

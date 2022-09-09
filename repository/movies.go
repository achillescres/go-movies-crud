package repository

import (
	"fmt"
	"go-movies-crud/models"
	"strconv"
)

type Movies interface {
	GetMovies() ([]models.Movie, error)
	GetMovie(id string) (*models.Movie, error)
	CreateMovie(movie models.Movie) error
	UpdateMovie(id string, movie models.Movie) (bool, error)
	DeleteMovie(id string) (bool, error)
}

type MoviesStore struct {
	movies []models.Movie
}

func NewMoviesStore() *MoviesStore {
	return new(MoviesStore)
}

func (store *MoviesStore) GetMovies() ([]models.Movie, error) {
	return store.movies, nil
}

func (store *MoviesStore) GetMovie(id string) (*models.Movie, error) {
	for _, movie := range store.movies {
		if movie.Id == id {
			return &movie, nil
		}
	}

	return nil, nil
}

func (store *MoviesStore) CreateMovie(movie models.Movie) error {
	fmt.Print(movie.Id)
	movie.Id = strconv.Itoa(len(store.movies))
	store.movies = append(store.movies, movie)
	return nil
}

func (store *MoviesStore) UpdateMovie(id string, newMovie models.Movie) (bool, error) {
	newMovie.Id = id
	for index, movie := range store.movies {
		if movie.Id == id {
			store.movies[index] = newMovie
			return true, nil
		}
	}

	return false, nil
}

func (store *MoviesStore) DeleteMovie(id string) (bool, error) {
	for index, movie := range store.movies {
		if movie.Id == id {
			store.movies = append(store.movies[:index], store.movies[index+1:]...)
			return true, nil
		}
	}

	return false, nil
}

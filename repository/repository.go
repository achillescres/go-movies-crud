package repository

type Repository struct {
	Movies Movies
}

func NewRepository() *Repository {
	return &Repository{
		Movies: NewMoviesStore(),
	}
}

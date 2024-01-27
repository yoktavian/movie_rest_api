package repository

import "movies/entity"

type MovieRepository interface {
	GetMovieByID(id string) (entity.Movie, error)
}

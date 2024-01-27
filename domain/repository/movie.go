package repository

import "movies/entity"

type MovieRepository interface {
	ReadByID(id string) (entity.Movie, error)
}

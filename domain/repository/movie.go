package repository

import "movies/entity"

type MovieRepository interface {
	Read(limit int, offset int) ([]entity.Movie, error)
	ReadByID(id string) (entity.Movie, error)
}

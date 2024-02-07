package repository

import "movies/entity"

type MovieRepository interface {
	Create(request entity.MovieRequest) (entity.Movie, error)
	Read(limit int, offset int) ([]entity.Movie, error)
	ReadByID(id string) (entity.Movie, error)
}

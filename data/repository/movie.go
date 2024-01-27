package repository

import (
	"movies/data/sql"
	"movies/domain/repository"
	"movies/entity"
)

type movieRepository struct {
	fakeSql sql.FakeSql
}

func NewMoveRepository(fakeSql sql.FakeSql) repository.MovieRepository {
	return &movieRepository{
		fakeSql: fakeSql,
	}
}

func (r *movieRepository) ReadByID(id string) (entity.Movie, error) {
	result := r.fakeSql.Get()
	var movie entity.Movie
	movie.FromMap(result)
	return movie, nil
}

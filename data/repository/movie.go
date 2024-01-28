package repository

import (
	"movies/domain/repository"
	"movies/entity"

	"gorm.io/gorm"
)

type movieRepository struct {
	dbe *gorm.DB
}

func NewMoveRepository(dbe *gorm.DB) repository.MovieRepository {
	return &movieRepository{dbe}
}

func (r *movieRepository) Read(limit int, offset int) ([]entity.Movie, error) {
	var movies []entity.Movie
	res := r.dbe.Table("movie").Select("id", "name", "link", "rating").Limit(limit).Offset(offset).Find(&movies)
	if res.Error != nil {
		return []entity.Movie{}, res.Error
	}

	return movies, nil
}

func (r *movieRepository) ReadByID(id string) (entity.Movie, error) {
	var movie entity.Movie
	res := r.dbe.Table("movie").Select("id", "name", "link", "rating").Where("id = ?", id).Find(&movie)
	if res.Error != nil {
		return entity.Movie{}, res.Error
	}

	return movie, nil
}

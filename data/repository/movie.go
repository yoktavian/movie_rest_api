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

func (r *movieRepository) Create(request entity.MovieRequest) (entity.Movie, error) {
	movie := entity.Movie{
		ID:     request.ID,
		Name:   request.Name,
		Link:   request.Link,
		Rating: request.Rating,
	}
	res := r.dbe.Table("movie").Create(&movie)
	if res.Error != nil {
		return entity.Movie{}, res.Error
	}

	return movie, nil
}

func (r *movieRepository) Read(limit int, offset int) ([]entity.Movie, error) {
	var movies []entity.Movie
	res := r.dbe.Table("movie").Select("movie.id", "movie.name", "movie.link", "movie.rating", "movie_creator.name", "movie_creator.company").Joins("left join movie_creator on movie_creator.id = movie.creator_id").Limit(limit).Offset(offset).Find(&movies)
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

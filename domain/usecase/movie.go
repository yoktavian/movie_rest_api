package usecase

import (
	"movies/domain/repository"
	"movies/entity"
)

type MovieUsecase interface {
	Read(limit int, offset int) ([]entity.Movie, error)
	ReadByID(id string) (entity.Movie, error)
}

type movieUsecase struct {
	MovieRepo repository.MovieRepository
}

func NewMovieUsecase(movieRepo repository.MovieRepository) MovieUsecase {
	return &movieUsecase{
		MovieRepo: movieRepo,
	}
}

func (u *movieUsecase) Read(limit int, offset int) ([]entity.Movie, error) {
	return u.MovieRepo.Read(limit, offset)
}

func (u *movieUsecase) ReadByID(id string) (entity.Movie, error) {
	return u.MovieRepo.ReadByID(id)
}

package usecase

import (
	"movies/domain/repository"
	"movies/entity"
)

type MovieUsecase interface {
	GetMovieByID(id string) (entity.Movie, error)
}

type movieUsecase struct {
	MovieRepo repository.MovieRepository
}

func NewMovieUsecase(movieRepo repository.MovieRepository) MovieUsecase {
	return &movieUsecase{
		MovieRepo: movieRepo,
	}
}

func (u *movieUsecase) GetMovieByID(id string) (entity.Movie, error) {
	return u.MovieRepo.GetMovieByID(id)
}

package usecase

import (
	"movies/domain/repository"
	"movies/entity"

	"github.com/go-playground/validator/v10"
)

type MovieUsecase interface {
	Create(request entity.MovieRequest) (entity.Movie, error)
	Read(limit int, offset int) ([]entity.Movie, error)
	ReadByID(id string) (entity.Movie, error)
}

type movieUsecase struct {
	MovieRepo repository.MovieRepository
	Validator *validator.Validate
}

func NewMovieUsecase(
	movieRepo repository.MovieRepository,
	validator *validator.Validate,
) MovieUsecase {
	return &movieUsecase{
		MovieRepo: movieRepo,
		Validator: validator,
	}
}

func (u *movieUsecase) Create(request entity.MovieRequest) (entity.Movie, error) {
	err := u.Validator.Struct(request)
	if err != nil {
		return entity.Movie{}, err
	}

	return u.MovieRepo.Create(request)
}

func (u *movieUsecase) Read(limit int, offset int) ([]entity.Movie, error) {
	return u.MovieRepo.Read(limit, offset)
}

func (u *movieUsecase) ReadByID(id string) (entity.Movie, error) {
	return u.MovieRepo.ReadByID(id)
}

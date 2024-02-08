package usecase

import (
	"fmt"
	"movies/entity"
)

type MovieHorrorUsecase struct {
	MovieUsecase
}

func (u MovieHorrorUsecase) Create(request entity.MovieRequest) (entity.Movie, error) {
	movieID := request.ID
	if movieID != "" {
		movieID = fmt.Sprintf("HRR%s", request.ID)
	}

	cmrRequest := entity.MovieRequest{
		ID:     movieID,
		Name:   request.Name,
		Link:   request.Link,
		Rating: request.Rating,
	}

	return u.MovieUsecase.Create(cmrRequest)
}

func NewMovieMovieHorrorUsecaseUsecase(base MovieUsecase) MovieUsecase {
	return &MovieHorrorUsecase{base}
}

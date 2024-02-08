package usecase

import (
	"fmt"
	"movies/entity"
)

type MovieRomanceComedyUsecase struct {
	MovieUsecase
}

func (u MovieRomanceComedyUsecase) Create(request entity.MovieRequest) (entity.Movie, error) {
	movieID := request.ID
	if movieID != "" {
		movieID = fmt.Sprintf("CMR%s", request.ID)
	}

	cmrRequest := entity.MovieRequest{
		ID:     movieID,
		Name:   request.Name,
		Link:   request.Link,
		Rating: request.Rating,
	}

	return u.MovieUsecase.Create(cmrRequest)
}

func NewMovieRomanceComedyUsecase(base MovieUsecase) MovieUsecase {
	return &MovieRomanceComedyUsecase{base}
}

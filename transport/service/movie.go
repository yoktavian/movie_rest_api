package service

import (
	"fmt"
	"movies/domain/usecase"
	"movies/entity"
	"movies/utils"

	"github.com/gin-gonic/gin"
)

type MovieService interface {
	Read(c *gin.Context)
	ReadByID(c *gin.Context)
}

type movieService struct {
	usecase usecase.MovieUsecase
}

func NewMovieService(usecase usecase.MovieUsecase) MovieService {
	return &movieService{usecase}
}

func (s *movieService) Read(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	fmt.Println(limit)
	fmt.Println(offset)
	movie, err := s.usecase.ReadByID("")
	if err != nil {
		utils.HandleError(c, 400, "error cuy")
		return
	}

	movies := []entity.Movie{
		movie,
		{
			ID:     "5",
			Name:   "Other movie",
			Link:   "other link",
			Rating: 1,
		},
	}

	utils.HandleSucces(c, movies)
}

func (s *movieService) ReadByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.HandleError(c, 401, "error cuy, id na eweuh")
		return
	}

	movie, err := s.usecase.ReadByID(id)
	if err != nil {
		utils.HandleError(c, 400, "error cuy")
		return
	}

	utils.HandleSucces(c, movie)
}

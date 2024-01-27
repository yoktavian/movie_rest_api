package handler

import (
	"fmt"
	"movies/domain/usecase"
	"movies/entity"
	"movies/interfaces/response"

	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	Read(c *gin.Context)
	ReadByID(c *gin.Context)
}

type movieHandler struct {
	usecase usecase.MovieUsecase
}

func NewMovieHandler(r *gin.Engine, u usecase.MovieUsecase) MovieHandler {
	return &movieHandler{u}
}

func (s *movieHandler) Read(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	fmt.Println(limit)
	fmt.Println(offset)
	movie, err := s.usecase.ReadByID("")
	if err != nil {
		response.Error(c, 400, "error cuy")
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

	response.Success(c, movies)
}

func (s *movieHandler) ReadByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, 401, "error cuy, id na eweuh")
		return
	}

	movie, err := s.usecase.ReadByID(id)
	if err != nil {
		response.Error(c, 400, "error cuy")
		return
	}

	response.Success(c, movie)
}

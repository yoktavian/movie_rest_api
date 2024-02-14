package handler

import (
	"movies/domain/usecase"
	"movies/entity"
	"movies/interfaces/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	ReadByID(c *gin.Context)
}

type movieHandler struct {
	usecases map[string]usecase.MovieUsecase
}

func NewMovieHandler(r *gin.Engine, u map[string]usecase.MovieUsecase) MovieHandler {
	return &movieHandler{u}
}

func (s *movieHandler) Create(c *gin.Context) {
	var movieRequest entity.MovieRequest

	if err := c.BindJSON(&movieRequest); err != nil {
		response.PublishCustomError(c, response.BadRequest, err.Error())
		return
	}

	// depends on type, need to differentiate the ID for each movie
	movie, err := s.usecases[movieRequest.Type].Create(movieRequest)

	if err != nil {
		response.PublishCustomError(c, response.BadRequest, err.Error())
		return
	}

	response.Publish(c, response.Successfull, movie)
}

func (s *movieHandler) Read(c *gin.Context) {
	movieType := c.Query("type")
	limitQuery := c.Query("limit")
	offsetQuery := c.Query("offset")
	limit, _ := strconv.Atoi(limitQuery)
	offset, _ := strconv.Atoi(offsetQuery)

	// use base, no need to differentiate based on movie type
	movies, err := s.usecases[movieType].Read(limit, offset)
	if err != nil {
		response.Publish(c, response.BadRequest, nil)
		return
	}

	response.Publish(c, response.Successfull, movies)
}

func (s *movieHandler) ReadByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Publish(c, response.BadRequest, nil)
		return
	}

	// use base, no need to differentiate based on movie type
	_, err := s.usecases[""].ReadByID(id)
	if err != nil {
		response.Publish(c, response.BadRequest, nil)
		return
	}

	response.Publish(c, 501, nil)
}

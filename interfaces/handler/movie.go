package handler

import (
	"fmt"
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

	movie, err := s.usecases[s.getMovieType(movieRequest.Type)].Create(movieRequest)

	if err != nil {
		response.PublishCustomError(c, response.BadRequest, err.Error())
		return
	}

	response.Publish(c, response.Successfull, movie)
}

func (s *movieHandler) Read(c *gin.Context) {
	limitQuery := c.Query("limit")
	offsetQuery := c.Query("offset")
	limit, err := strconv.Atoi(limitQuery)
	offset, err := strconv.Atoi(offsetQuery)

	movies, err := s.usecases[s.getMovieType("cmr")].Read(limit, offset)
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

	_, err := s.usecases[s.getMovieType("")].ReadByID(id)
	if err != nil {
		response.Publish(c, response.BadRequest, nil)
		return
	}

	response.Publish(c, 501, nil)
}

func (s *movieHandler) getMovieType(payloadMovieType string) string {
	if payloadMovieType == "" {
		return "df"
	}

	return payloadMovieType
}

package handler

import (
	"movies/transport/service"

	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	RegisterRoute(r *gin.Engine)
}

type movieHandler struct {
	service service.MovieService
}

func NewMovieHandler(r *gin.Engine, s service.MovieService) {
	handler := movieHandler{s}
	handler.RegisterRoute(r)
}

func (h *movieHandler) RegisterRoute(r *gin.Engine) {
	r.GET("/movie", h.service.Read)
	r.GET("/movie/:id", h.service.ReadByID)
}

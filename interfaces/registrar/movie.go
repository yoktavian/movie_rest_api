package registrar

import (
	"movies/data/repository"
	"movies/domain/usecase"
	"movies/interfaces/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type movieRegistrar struct {
	g   *gin.Engine
	dbe *gorm.DB
}

func NewMovieRegistrar(g *gin.Engine, dbe *gorm.DB) Registrar {
	return &movieRegistrar{g, dbe}
}

func (r *movieRegistrar) Register() {
	movieRepo := repository.NewMoveRepository(r.dbe)
	movieUsecase := usecase.NewMovieUsecase(movieRepo)

	h := handler.NewMovieHandler(r.g, movieUsecase)
	r.g.GET("/movie", h.Read)
	r.g.GET("/movie/:id", h.ReadByID)
}

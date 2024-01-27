package registrar

import (
	"movies/data/repository"
	"movies/data/sql"
	"movies/domain/usecase"
	"movies/interfaces/handler"

	"github.com/gin-gonic/gin"
)

type movieRegistrar struct {
	g       *gin.Engine
	fakeSql sql.FakeSql
}

func NewMovieRegistrar(g *gin.Engine, fakeSql sql.FakeSql) Registrar {
	return &movieRegistrar{g, fakeSql}
}

func (r *movieRegistrar) Register() {
	movieRepo := repository.NewMoveRepository(r.fakeSql)
	movieUsecase := usecase.NewMovieUsecase(movieRepo)

	h := handler.NewMovieHandler(r.g, movieUsecase)
	r.g.GET("/movie", h.Read)
	r.g.GET("/movie/:id", h.ReadByID)
}

package registrar

import (
	"movies/data/repository"
	"movies/data/sql"
	"movies/domain/usecase"
	"movies/transport/handler"
	"movies/transport/service"

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
	movieService := service.NewMovieService(movieUsecase)
	handler.NewMovieHandler(r.g, movieService)
}

package registrar

import (
	"movies/data/repository"
	"movies/domain/usecase"
	"movies/interfaces/handler"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	validator := validator.New(validator.WithRequiredStructEnabled())

	movieRepo := repository.NewMoveRepository(r.dbe)
	baseMovieUsecase := usecase.NewBaseMovieUsecase(movieRepo, validator)
	comedyMovieUsecase := usecase.NewMovieRomanceComedyUsecase(baseMovieUsecase)
	horrorMovieUsecase := usecase.NewMovieMovieHorrorUsecaseUsecase(baseMovieUsecase)

	usecases := map[string]usecase.MovieUsecase{
		"comedy": comedyMovieUsecase,
		"horror": horrorMovieUsecase,
	}

	h := handler.NewMovieHandler(r.g, usecases)
	r.g.GET("/movie", h.Read)
	r.g.GET("/movie/:id", h.ReadByID)
	r.g.POST("/movie", h.Create)
}

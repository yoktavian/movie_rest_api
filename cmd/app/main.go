package main

import (
	"fmt"
	"movies/data/repository"
	"movies/data/sql"
	"movies/domain/usecase"
)

func main() {
	fakeSql := sql.NewFakeSql()
	movieRepo := repository.NewMoveRepository(fakeSql)
	movieUsecase := usecase.MovieUsecase(movieRepo)
	res, err := movieUsecase.GetMovieByID("1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.ID)
		fmt.Println(res.Name)
		fmt.Println(res.Link)
		fmt.Println(res.Rating)
	}
}

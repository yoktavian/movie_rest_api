package main

import (
	"log"
	"movies/data/sql"
	"movies/transport/registrar"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	fakeSql := sql.NewFakeSql()

	registrarGroup := []registrar.Registrar{
		registrar.NewMovieRegistrar(g, fakeSql),
		// registrar.NewMovieRegistrar(g, fakeSql),
	}
	registerAllRegistrar(registrarGroup)

	err := g.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func registerAllRegistrar(registrarGroup []registrar.Registrar) {
	for i := 0; i < len(registrarGroup); i++ {
		registrarGroup[i].Register()
	}
}

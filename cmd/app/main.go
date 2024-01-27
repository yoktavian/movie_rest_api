package main

import (
	"log"
	"movies/data/sql"
	"movies/interfaces/registrar"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	// registrar
	rg := registrarGroup(g)
	registerAllRegistrar(rg)
	// server
	runServer(g)
}

func registrarGroup(g *gin.Engine) []registrar.Registrar {
	fakeSql := sql.NewFakeSql()

	return []registrar.Registrar{
		registrar.NewMovieRegistrar(g, fakeSql),
	}
}

func registerAllRegistrar(registrarGroup []registrar.Registrar) {
	for i := 0; i < len(registrarGroup); i++ {
		registrarGroup[i].Register()
	}
}

func runServer(g *gin.Engine) {
	err := g.Run()
	if err != nil {
		log.Fatal(err)
	}
}

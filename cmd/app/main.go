package main

import (
	"log"
	"movies/data/db"
	"movies/interfaces/registrar"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	g := gin.Default()
	// connect db
	dbe, _ := db.ConnectDB()
	// registrar
	rg := registrarGroup(g, dbe)
	registerAllRegistrar(rg)
	// server
	runServer(g)
}

func registrarGroup(g *gin.Engine, dbe *gorm.DB) []registrar.Registrar {
	return []registrar.Registrar{
		registrar.NewMovieRegistrar(g, dbe),
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


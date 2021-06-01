package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terrytay/go-backend/types"
	"log"
	"net/http"
)

type App struct {
	router *gin.Engine
}

func Initialize(controllers *[]types.Controller) *App {
	app := App{router: gin.Default()}

	app.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	initializeControllers(app.router, controllers)

	return &app
}

func initializeControllers(router *gin.Engine, controllers *[]types.Controller) {
	if controllers != nil {
		for _, controller := range *controllers {
			controller.InitializeRoutes(router)
		}
	}
}

func (a *App) Start(port string) {
	err := a.router.Run(port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Running on port: %s", port)
}
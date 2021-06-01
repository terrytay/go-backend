package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terrytay/go-backend/types"
	"io"
	"log"
	"net/http"
	"os"
)

type App struct {
	router *gin.Engine
}

func Initialize(controllers *[]types.Controller) *App {

	// Set up gin logger
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// gin app
	app := App{router: gin.Default()}

	app.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	initializeMiddlewares()

	initializeControllers(app.router, controllers)

	return &app
}

func initializeMiddlewares() {

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
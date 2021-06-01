package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terrytay/go-backend/helpers"
	"github.com/terrytay/go-backend/middlwares"
	"github.com/terrytay/go-backend/types"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type App struct {
	router *gin.Engine
}

func Initialize(controllers *[]types.Controller) *App {

	// Set up gin logger
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Non-debug mode
	//gin.SetMode(gin.ReleaseMode)


	// gin app
	app := App{router: gin.New()}

	// TODO remove this when authentication done

	app.router.GET("/", func(c *gin.Context) {
		i, _ := strconv.ParseUint("test", 10, 64)
		token, err := helpers.CreateToken(i)
		if err != nil {
			log.Fatal("unable to generate token")
		}

		c.JSON(http.StatusOK, gin.H{"token": token} )
	})

	initializeMiddlewares(app.router)

	initializeControllers(app.router, controllers)

	return &app
}

func initializeMiddlewares(router *gin.Engine) {
	router.Use(middlwares.TokenAuthMiddleware())
}

func initializeControllers(router *gin.Engine, controllers *[]types.Controller) {
	if controllers != nil {
		for _, controller := range *controllers {
			controller.InitializeRoutes(router)
		}
	}
}

func (a *App) Start(port string) {

	if helpers.GetEnvVariable("ENV") == "development" {
		log.Printf("server running on http://localhost%s", port)
	}

	err := a.router.Run(port)
	if err != nil {
		log.Fatal(err)
	}

}
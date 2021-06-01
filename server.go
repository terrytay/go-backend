package main

import (
	"github.com/terrytay/go-backend/controllers"
	"github.com/terrytay/go-backend/db"
	"github.com/terrytay/go-backend/helpers"
	"github.com/terrytay/go-backend/repositories"
	"github.com/terrytay/go-backend/services"
	"github.com/terrytay/go-backend/types"
)

type Repositories struct {
	userRepository repositories.IUserRepository
}

type Services struct {
	userService services.IUserService
}

func getRepositories(dbClient *db.Client) *Repositories {
	return &Repositories{
		userRepository: &repositories.UserRepository{DbClient: dbClient},
	}

}

func getServices(repos *Repositories) *Services {
	return &Services{
		userService: &services.UserService{Repo: repos.userRepository},
	}
}

func getControllers(services *Services) *[]types.Controller {
	return &[]types.Controller{
		&controllers.HealthController{},
		&controllers.UserController{UserService: services.userService},
	}
}

func initApp(dbClient *db.Client) {
	repos := getRepositories(dbClient)
	svcs := getServices(repos)
	controllers := getControllers(svcs)

	app := Initialize(controllers)

	port := helpers.GetEnvVariable("PORT")

	if port == "" {
		port = ":3000"
	}

	app.Start(port)

}
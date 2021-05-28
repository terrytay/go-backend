package main

import (
	"fmt"
	"github.com/terrytay/go-backend/db"
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

func getControllers(services *Services) []types.Controller {
	return nil
}

func initApp(dbClient *db.Client) {
	repos := getRepositories(dbClient)
	svcs := getServices(repos)
	//controllers := getControllers(svcs)
	fmt.Println(svcs)
}
package services

import (
	"github.com/terrytay/backend-architecture/entities"
	"github.com/terrytay/backend-architecture/repositories"
)

type IUserService interface {
	GetUser(username string) *entities.User
}

type UserService struct {
	Repo repositories.IUserRepository
}

func (s *UserService) GetUser(username string) *entities.User {
	return s.Repo.GetUser(username)
}
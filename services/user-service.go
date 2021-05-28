package services

import (
	"github.com/terrytay/go-backend/entities"
	"github.com/terrytay/go-backend/repositories"
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
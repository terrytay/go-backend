package services

import (
	"github.com/terrytay/go-backend/entities"
	"github.com/terrytay/go-backend/repositories"
)

type IUserService interface {
	GetUser(username string) (*entities.User, error)
}

type UserService struct {
	Repo repositories.IUserRepository
}

func (s *UserService) GetUser(username string) (*entities.User, error) {
	user, err := s.Repo.GetUser(username)

	if err != nil {
		return nil, err
	}

	return user, nil
}
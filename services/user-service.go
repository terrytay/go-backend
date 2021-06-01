package services

import (
	"github.com/gin-gonic/gin"
	"github.com/terrytay/go-backend/entities"
	"github.com/terrytay/go-backend/repositories"
)

type IUserService interface {
	GetUser(c *gin.Context, username string) (*entities.User, error)
}

type UserService struct {
	Repo repositories.IUserRepository
}

func (s *UserService) GetUser(c *gin.Context, username string) (*entities.User, error) {
	user, err := s.Repo.GetUser(c, username)

	if err != nil {
		return nil, err
	}

	return user, nil
}
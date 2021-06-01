package repositories

import (
	"errors"
	"github.com/terrytay/go-backend/db"
	"github.com/terrytay/go-backend/entities"
)

type IUserRepository interface {
	GetUser(username string) (*entities.User, error)
}

type UserRepository struct {
	DbClient *db.Client
}

func (r *UserRepository) GetUser(username string) (*entities.User, error) {
	if username == "fake" {
		return nil, errors.New("user not found")
	}
	return &entities.User{
		Username: username,
		Password: "testing",
	}, nil
}
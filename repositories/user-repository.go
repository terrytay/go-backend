package repositories

import (
	"github.com/terrytay/go-backend/db"
	"github.com/terrytay/go-backend/entities"
)

type IUserRepository interface {
	GetUser(username string) *entities.User
}

type UserRepository struct {
	DbClient *db.Client
}

func (r *UserRepository) GetUser(username string) *entities.User {
	return &entities.User{
		Username: username,
		PasswordHash: "testing",
	}
}
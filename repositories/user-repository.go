package repositories

import (
	"github.com/terrytay/backend-architecture/db"
	"github.com/terrytay/backend-architecture/entities"
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
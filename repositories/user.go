package repositories

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/terrytay/go-backend/db"
	"github.com/terrytay/go-backend/entities"
)

type IUserRepository interface {
	GetUser(c *gin.Context, username string) (*entities.User, error)
}

type UserRepository struct {
	DbClient *db.Client
}

func (r *UserRepository) GetUser(c *gin.Context, username string) (*entities.User, error) {
	user := new(entities.UserDTO)

	err := r.DbClient.GetConnection().ModelContext(c, user).Where("username = ?", username).Select()
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user.ToEntity(), nil
}
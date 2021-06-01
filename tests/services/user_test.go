package test_controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/terrytay/go-backend/entities"
	"github.com/terrytay/go-backend/services"
	"testing"
)

type mockUserRepo struct {
}

func (s mockUserRepo) GetUser(c *gin.Context, username string) (*entities.User, error) {
	if username != "admin" {
		return nil, errors.New("user not found")
	}

	return &entities.User{
		Username: "admin",
		Password: "123",
	}, nil
}

func TestCannotFindUser(t *testing.T) {
	userService := services.UserService{Repo: mockUserRepo{}}
	_, err := userService.GetUser(nil, "test")
	if err == nil {
		t.Errorf("invalid user does not throw error")
	}
}

func TestSuccessUser(t *testing.T) {
	userService := services.UserService{Repo: mockUserRepo{}}
	_, err := userService.GetUser(nil, "admin")
	if err != nil {
		t.Errorf("does not return valid user")
	}
}
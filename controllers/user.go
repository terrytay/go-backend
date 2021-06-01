package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/terrytay/go-backend/services"
	"net/http"
)

type UserController struct {
	UserService services.IUserService
}

func (controller *UserController) InitializeRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/:username", controller.get)
	}
}

func (controller *UserController) get(c *gin.Context) {
	username := c.Param("username")

	user, err := controller.UserService.GetUser(c, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}
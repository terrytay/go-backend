package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HealthController struct{
}

func (controller *HealthController) InitializeRoutes(router *gin.Engine) {
	router.GET("/health", controller.get)
}

func (controller *HealthController) get(c *gin.Context) {
	date := time.Now()
	c.JSON(http.StatusOK, gin.H{"success": true, "timestamp": date, "message": "OK"})

}
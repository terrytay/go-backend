package middlwares

import (
	"github.com/gin-gonic/gin"
	"github.com/terrytay/go-backend/helpers"
	"github.com/terrytay/go-backend/types"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, types.GenericResponse{
				Success: false,
				Message: "invalid credentials",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
package middlewares

import (
	"github.com/gin-gonic/gin"
	"manage-jwt/auth"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}

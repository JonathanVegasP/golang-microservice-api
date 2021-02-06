package middlewares

import (
	"flutter-store-api/infrastructure/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isAuthenticated := auth.ValidateToken(c.Request); !isAuthenticated {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}

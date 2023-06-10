package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kykurniawan/go-jwt-auth/app/services"
	"github.com/kykurniawan/go-jwt-auth/custom_errors"
)

func AuthenticationMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(custom_errors.NewUnauthorizedError("authorization header is required"))
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			c.Error(custom_errors.NewUnauthorizedError(err.Error()))
			c.Abort()
			return
		}

		tokenType := token.Claims.(jwt.MapClaims)["typ"].(string)

		if tokenType != "access_token" {
			c.Error(custom_errors.NewUnauthorizedError("invalid token type"))
			c.Abort()
			return
		}

		userId := token.Claims.(jwt.MapClaims)["id"].(float64)

		c.Set("userId", userId)

		c.Next()
	}
}

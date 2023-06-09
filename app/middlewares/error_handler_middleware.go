package middlewares

import (
	"github.com/kykurniawan/go-jwt-auth/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()

		if err != nil {
			switch e := err.Err.(type) {
			case *app.NotFoundError:
				c.JSON(e.Code, gin.H{
					"message": "not found error",
					"data":    nil,
					"error":   e.Error(),
				})
			case *app.BadRequestError:
				c.JSON(e.Code, gin.H{
					"message": "bad request error",
					"data":    nil,
					"error":   e.Error(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "internal server error",
					"data":    nil,
					"error":   err.Error(),
				})
			}
		}
	}
}

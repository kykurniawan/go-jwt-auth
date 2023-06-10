package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kykurniawan/go-jwt-auth/custom_errors"
	"github.com/kykurniawan/go-jwt-auth/helpers"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()

		if err != nil {
			switch e := err.Err.(type) {
			case *custom_errors.NotFoundError:
				c.JSON(e.Code, helpers.FormatResponse("not found", nil, e.Error()))
			case *custom_errors.BadRequestError:
				c.JSON(e.Code, helpers.FormatResponse("bad request", nil, e.Error()))
			case *custom_errors.ValidationError:
				c.JSON(e.Code, helpers.FormatResponse(e.Error(), nil, gin.H{
					"fields": e.Fields,
					"old":    e.OldData,
				}))
			default:
				c.JSON(http.StatusInternalServerError, helpers.FormatResponse("internal server error", nil, e.Error()))
			}
		}
	}
}

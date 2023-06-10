package main

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kykurniawan/go-jwt-auth/app/middlewares"
	"github.com/kykurniawan/go-jwt-auth/custom_validators"
	"github.com/kykurniawan/go-jwt-auth/database"
	"github.com/kykurniawan/go-jwt-auth/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func init() {
	database.CreateConnection()
}

func main() {
	app := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.SetTagName("validate")
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		v.RegisterValidation("unique_email", custom_validators.UniqueEmailValidator)
	}

	app.SetTrustedProxies([]string{"127.0.0.1"})

	app.Use(middlewares.ErrorHandlerMiddleware())

	routes.RegisterApiRoutes(app)

	app.Run(":3000")
}

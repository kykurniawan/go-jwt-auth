package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kykurniawan/go-jwt-auth/app/middlewares"
	"github.com/kykurniawan/go-jwt-auth/configs"
	"github.com/kykurniawan/go-jwt-auth/custom_validators"
	"github.com/kykurniawan/go-jwt-auth/database"
	"github.com/kykurniawan/go-jwt-auth/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.CreateConnection()
}

func main() {
	gin.SetMode(configs.App().Mode)
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())

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

	app.SetTrustedProxies(configs.App().TrustedProxies)

	app.Use(middlewares.ErrorHandlerMiddleware())

	routes.RegisterApiRoutes(app)

	app.Run(fmt.Sprintf("%s:%s", configs.App().Host, configs.App().Port))
}

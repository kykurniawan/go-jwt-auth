package main

import (
	"github.com/kykurniawan/go-jwt-auth/app/middlewares"
	"github.com/kykurniawan/go-jwt-auth/database"
	"github.com/kykurniawan/go-jwt-auth/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	database.CreateConnection()
}

func main() {
	app := gin.Default()

	app.SetTrustedProxies([]string{"127.0.0.1"})

	app.Use(middlewares.ErrorHandlerMiddleware())

	routes.RegisterApiRoutes(app)

	app.Run(":3000")
}

package routes

import (
	"github.com/kykurniawan/go-jwt-auth/app/controllers"
	"github.com/kykurniawan/go-jwt-auth/app/repositories"
	"github.com/kykurniawan/go-jwt-auth/database"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(app *gin.Engine) {
	userRepository := repositories.NewUserRepository(database.Connection)

	pageController := controllers.NewPageController()
	userController := controllers.NewUserController(userRepository)

	router := app.Group("/api")
	{
		router.GET("/", pageController.Home)

		router.GET("/users", userController.Index)
		router.POST("/users", userController.Store)
		router.GET("/users/:id", userController.Show)
		router.PUT("/users/:id", userController.Update)
		router.DELETE("/users/:id", userController.Destroy)
	}
}

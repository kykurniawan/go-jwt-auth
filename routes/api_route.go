package routes

import (
	"github.com/kykurniawan/go-jwt-auth/app/controllers"
	"github.com/kykurniawan/go-jwt-auth/app/middlewares"
	"github.com/kykurniawan/go-jwt-auth/app/repositories"
	"github.com/kykurniawan/go-jwt-auth/app/services"
	"github.com/kykurniawan/go-jwt-auth/database"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(app *gin.Engine) {
	userRepository := repositories.NewUserRepository(database.Connection)
	userSessionRepository := repositories.NewUserSessionRepository(database.Connection)

	authService := services.NewAuthService(userRepository, userSessionRepository)
	userService := services.NewUserService(userRepository)

	pageController := controllers.NewPageController()
	userController := controllers.NewUserController(userRepository)
	authController := controllers.NewAuthController(authService, userService)

	router := app.Group("/api")
	{
		router.GET("/", pageController.Home)
		router.POST("/login", authController.Login)
		router.POST("/register", authController.Register)
		router.POST("/logout", middlewares.AuthenticationMiddleware(authService), authController.Logout)
		router.POST("/refresh-token", authController.RefreshToken)

		router = router.Group("/")
		router.Use(middlewares.AuthenticationMiddleware(authService))
		{
			router.GET("/users", userController.Index)
			router.POST("/users", userController.Store)
			router.GET("/users/:id", userController.Show)
			router.PUT("/users/:id", userController.Update)
			router.DELETE("/users/:id", userController.Destroy)
		}
	}
}

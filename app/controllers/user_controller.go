package controllers

import (
	"net/http"
	"strconv"

	"github.com/kykurniawan/go-jwt-auth/app"
	"github.com/kykurniawan/go-jwt-auth/app/models"
	"github.com/kykurniawan/go-jwt-auth/app/repositories"
	"github.com/kykurniawan/go-jwt-auth/helpers"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository *repositories.UserRepository
}

func NewUserController(userRepository *repositories.UserRepository) *UserController {
	return &UserController{userRepository}
}

func (controller *UserController) Index(c *gin.Context) {
	users, err := controller.userRepository.FindAll()

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    users,
		"error":   nil,
	})
}

func (controller *UserController) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Error(err)
		return
	}

	user, err := controller.userRepository.FindById(uint(id))

	if err != nil {
		c.Error(app.NewNotFoundError("User not found"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    user,
		"error":   nil,
	})
}

func (controller *UserController) Store(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.Error(app.NewBadRequestError("Invalid request body"))
		return
	}

	hash, hashErr := helpers.HashPassword(user.Password)

	if hashErr != nil {
		c.Error(hashErr)
		return
	}

	user.Password = string(hash)

	err = controller.userRepository.Create(&user)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    user,
		"error":   nil,
	})
}

func (controller *UserController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    "updating an user",
		"error":   nil,
	})
}

func (controller *UserController) Destroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    "deleting an user",
		"error":   nil,
	})
}

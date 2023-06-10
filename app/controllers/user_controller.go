package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kykurniawan/go-jwt-auth/app/models"
	"github.com/kykurniawan/go-jwt-auth/app/repositories"
	"github.com/kykurniawan/go-jwt-auth/app/requests"
	"github.com/kykurniawan/go-jwt-auth/custom_errors"
	"github.com/kykurniawan/go-jwt-auth/helpers"
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

	c.JSON(http.StatusOK, helpers.FormatResponse("ok", users, nil))
}

func (controller *UserController) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Error(err)
		return
	}

	user, err := controller.userRepository.FindById(uint(id))

	if err != nil {
		c.Error(custom_errors.NewNotFoundError("User not found"))
		return
	}

	c.JSON(http.StatusOK, helpers.FormatResponse("ok", user, nil))
}

func (controller *UserController) Store(c *gin.Context) {
	var request requests.CreateUserRequest
	var user models.User

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.Error(custom_errors.NewValidationError("Validation error", err.(validator.ValidationErrors), request))
		return
	}

	hash, hashErr := helpers.HashPassword(user.Password)

	if hashErr != nil {
		c.Error(hashErr)
		return
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hash)

	err = controller.userRepository.Create(&user)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, helpers.FormatResponse("created", user, nil))
}

func (controller *UserController) Update(c *gin.Context) {
	var request requests.UpdateUserRequest

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Error(err)
		return
	}

	user, err := controller.userRepository.FindById(uint(id))

	if err != nil {
		c.Error(custom_errors.NewNotFoundError("User not found"))
		return
	}

	request.ID = uint(id)

	err = c.ShouldBindJSON(&request)

	if err != nil {
		c.Error(custom_errors.NewValidationError("Validation error", err.(validator.ValidationErrors), request))
		return
	}

	user.Name = request.Name
	user.Email = request.Email

	err = controller.userRepository.Update(user)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.FormatResponse("updated", user, nil))
}

func (controller *UserController) Destroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Error(err)
		return
	}

	user, err := controller.userRepository.FindById(uint(id))

	if err != nil {
		c.Error(custom_errors.NewNotFoundError("User not found"))
		return
	}

	err = controller.userRepository.Delete(user)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.FormatResponse("deleted", id, nil))
}

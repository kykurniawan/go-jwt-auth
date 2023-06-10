package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kykurniawan/go-jwt-auth/app/requests"
	"github.com/kykurniawan/go-jwt-auth/app/services"
	"github.com/kykurniawan/go-jwt-auth/custom_errors"
	"github.com/kykurniawan/go-jwt-auth/helpers"
)

type AuthCotroller struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthCotroller {
	return &AuthCotroller{
		authService,
	}
}

func (controller *AuthCotroller) Login(c *gin.Context) {
	var request requests.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(custom_errors.NewValidationError("validation error", err.(validator.ValidationErrors), request))
		return
	}

	token, err := controller.authService.Attempt(request.Email, request.Password)

	if err != nil {
		c.Error(custom_errors.NewUnauthorizedError("invalid credentials"))
		return
	}

	c.JSON(http.StatusOK, helpers.FormatResponse("login success", token, nil))
}

func (controller *AuthCotroller) Register(c *gin.Context) {
	//
}

func (controller *AuthCotroller) Logout(c *gin.Context) {
	var request requests.LogoutRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(custom_errors.NewValidationError("validation error", err.(validator.ValidationErrors), request))
		return
	}

	err := controller.authService.Logout(request.RefreshToken)

	if err != nil {
		c.Error(custom_errors.NewUnauthorizedError("invalid token"))
		return
	}

	c.JSON(http.StatusOK, helpers.FormatResponse("logout success", nil, nil))
}

func (controller *AuthCotroller) RefreshToken(c *gin.Context) {
	var request requests.RefreshTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(custom_errors.NewValidationError("validation error", err.(validator.ValidationErrors), request))
		return
	}

	token, err := controller.authService.Refresh(request.RefreshToken)

	if err != nil {
		c.Error(custom_errors.NewUnauthorizedError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.FormatResponse("refresh token success", token, nil))
}

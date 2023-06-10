package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

func Validate() *validator.Validate {
	return binding.Validator.Engine().(*validator.Validate)
}

func FormatResponse(message string, data interface{}, err interface{}) gin.H {
	return gin.H{
		"message": message,
		"data":    data,
		"error":   err,
	}
}

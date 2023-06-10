package helpers

import (
	"os"
	"path/filepath"
	"strconv"

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

func EnvString(key string, defaultValue string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return env
}

func EnvInt(key string, defaultValue int) int {
	env, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	res, err := strconv.Atoi(env)

	if err != nil {
		return defaultValue
	}

	return res
}

func EnvBool(key string, defaultValue bool) bool {
	env, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	res, err := strconv.ParseBool(env)

	if err != nil {
		return defaultValue
	}

	return res
}

func ExecutablePath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}

	return filepath.Dir(exePath)
}

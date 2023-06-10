package custom_validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/kykurniawan/go-jwt-auth/app/repositories"
	"github.com/kykurniawan/go-jwt-auth/database"
)

func UniqueEmailValidator(fieldLevel validator.FieldLevel) bool {
	var ignoreId uint
	email := fieldLevel.Field().String()
	if fieldLevel.Param() != "" {
		ignoreId = fieldLevel.Parent().FieldByName(fieldLevel.Param()).Interface().(uint)
	}

	if email == "" {
		return true
	}

	userRepository := repositories.NewUserRepository(database.Connection)

	return !userRepository.IsEmailExists(email, ignoreId)
}

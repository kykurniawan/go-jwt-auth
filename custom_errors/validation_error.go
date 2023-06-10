package custom_errors

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type customFieldError struct {
	Name    string      `json:"name"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

type ValidationError struct {
	Code    int
	Message string
	Fields  []customFieldError
	OldData interface{}
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(message string, validationErrors validator.ValidationErrors, old interface{}) *ValidationError {
	var fields []customFieldError
	for _, err := range validationErrors {
		fields = append(fields, formatValidatioError(err))
	}

	return &ValidationError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
		Fields:  fields,
		OldData: old,
	}
}

func formatValidatioError(err validator.FieldError) customFieldError {
	error := customFieldError{
		Name:  err.Field(),
		Value: err.Value(),
	}
	switch err.Tag() {
	case "required":
		error.Message = "This field is required"
	case "email":
		error.Message = "This field must be a valid email address"
	case "min":
		error.Message = "This field must be at least 6 characters"
	case "unique_email":
		error.Message = "This email has already been taken"
	default:
		error.Message = "This field is invalid"
	}

	return error
}

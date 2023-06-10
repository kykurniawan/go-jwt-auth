package custom_errors

import "net/http"

type BadRequestError struct {
	Code    int
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

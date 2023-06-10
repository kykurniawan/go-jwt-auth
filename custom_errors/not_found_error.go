package custom_errors

import "net/http"

type NotFoundError struct {
	Code    int
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

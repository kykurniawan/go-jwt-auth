package custom_errors

import "net/http"

type UnauthorizedError struct {
	Code    int
	Message string
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

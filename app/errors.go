package app

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

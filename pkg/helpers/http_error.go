package helpers

import "net/http"

type HttpError struct {
	Code    int
	Message string
}

func (httpError HttpError) Error() string {
	return httpError.Message
}

func NewBadRequestError(msg string) HttpError {
	return HttpError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func NewInternalServerError(msg string) HttpError {
	return HttpError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func NewNotFoundError(msg string) HttpError {
	return HttpError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

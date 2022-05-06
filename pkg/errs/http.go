package errs

import (
	"net/http"
)

const (
	ErrBadRequest          = "bad_request"
	ErrNotFound            = "not_found"
	ErrInternalServerError = "internal_server_error"
	ErrUnauthorized        = "unauthorized"
)

type Err struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"status"`
	Type       string      `json:"type"`
	Data       interface{} `json:"data"`
}

func newErr(statusCode int, message, errorType string, data interface{}) *Err {
	return &Err{
		message,
		statusCode,
		errorType,
		data,
	}
}

func (he *Err) Error() string {
	return he.Message
}

func NewBadRequestErr(message string, data interface{}) *Err {
	return newErr(http.StatusBadRequest, message, ErrBadRequest, data)
}

func NewNotFoundErr(message string, data interface{}) *Err {
	return newErr(http.StatusNotFound, message, ErrNotFound, data)
}

func NewInternalServerErr(message string, data interface{}) *Err {
	return newErr(http.StatusInternalServerError, message, ErrInternalServerError, data)
}

func NewUnauthorizedErr(message string, data interface{}) *Err {
	return newErr(http.StatusUnauthorized, message, ErrUnauthorized, data)
}

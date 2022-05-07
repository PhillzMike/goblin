package errs

import (
	"fmt"
	"net/http"
	"reflect"
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

func (e *Err) Error() string {
	return e.Message
}

func (e *Err) ErrorDetails() string {
	return fmt.Sprintf("msg: %s, statuscode: %d, type: %s, data: %+v", e.Message, e.StatusCode, e.Type, e.Data)
}

func (e *Err) Equals(err *Err) bool {
	if e.Message != err.Message || e.StatusCode != err.StatusCode ||
		e.Type != err.Type || !reflect.DeepEqual(e.Data, err.Data) {
		return false
	}
	return true
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

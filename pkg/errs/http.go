package errs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ErrBadRequest          = "bad_request"
	ErrNotFound            = "not_found"
	ErrInternalServerError = "internal_server_error"
	ErrUnauthorized        = "unauthorized"
)

// MarshalableErr is a custom type for a slice of error types
// It allows the error type to be 'marshable' by JSON.
type MarshalableErr []error

func (m MarshalableErr) MarshalJSON() ([]byte, error) {
	data := []byte("[")
	for i, err := range m {
		if i != 0 {
			data = append(data, ',')
		}
		j, err := json.Marshal(err.Error())
		if err != nil {
			return nil, err
		}
		data = append(data, j...)
	}
	data = append(data, ']')
	return data, nil
}

type Err struct {
	Message    string         `json:"message"`
	StatusCode int            `json:"status"`
	Type       string         `json:"type"`
	Data       MarshalableErr `json:"data"`
}

func NewErr(statusCode int, message, errorType string, data error) *Err {
	err := &Err{
		Message:    message,
		StatusCode: statusCode,
		Type:       errorType,
	}
	err.Add(data)
	return err
}

func (e *Err) GetStatusCode() int {
	return e.StatusCode
}

func (e *Err) GetType() string {
	return e.Type
}

func (e *Err) GetData() []error {
	return e.Data
}

func (e *Err) HasData() bool {
	return len(e.Data) > 0
}

func (e *Err) SetMessage(msg string) {
	e.Message = msg
}

func (e *Err) Add(data error) {
	if data != nil {
		e.Data = append(e.Data, data)
	}
}

// ErrorDetails returns a stringified representation of the
// fields contained in Err
func (e *Err) ErrorDetails() string {
	return fmt.Sprintf("msg: %s, statuscode: %d, type: %s, data: %+v", e.Message, e.StatusCode, e.Type, e.Data)
}

func (e *Err) Equals(err *Err) bool {
	if e.Message != err.Message || e.StatusCode != err.StatusCode ||
		e.Type != err.Type || !compareErrors(e.Data, err.Data) {
		return false
	}
	return true
}

// NewBadRequestErr returns the Err object for http bad_request
func NewBadRequestErr(message string, data error) *Err {
	return NewErr(http.StatusBadRequest, message, ErrBadRequest, data)
}

// NewNotFoundErr returns the Err object for http bad_request
func NewNotFoundErr(message string, data error) *Err {
	return NewErr(http.StatusNotFound, message, ErrNotFound, data)
}

// NewInternalServerErr returns the Err object for http bad_request
func NewInternalServerErr(message string, data error) *Err {
	return NewErr(http.StatusInternalServerError, message, ErrInternalServerError, data)
}

// NewUnauthorizedErr returns the Err object for http bad_request
func NewUnauthorizedErr(message string, data error) *Err {
	return NewErr(http.StatusUnauthorized, message, ErrUnauthorized, data)
}

func compareErrors(a, b []error) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Error() != b[i].Error() {
			return false
		}
	}
	return true
}

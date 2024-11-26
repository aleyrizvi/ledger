package engine

import (
	"net/http"
	"strings"
)

type APIError struct {
	Code     int      `json:"-"`
	Messages []string `json:"messages,omitempty"`
}

func (a APIError) Error() string {
	return strings.Join(a.Messages, "\n")
}

// ErrBadRequest returns an APIError with status code 400.
func ErrBadRequest(errors []string) APIError {
	return APIError{
		Code:     http.StatusBadRequest,
		Messages: errors,
	}
}

// ErrNotFound returns an APIError with status code 401.
func ErrNotFound() APIError {
	return APIError{
		Code:     http.StatusNotFound,
		Messages: nil,
	}
}

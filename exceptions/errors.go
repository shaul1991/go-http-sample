package exceptions

import (
	"fmt"
	"net/http"
)

// AppError represents a custom application error
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, error: %v", e.Code, e.Message, e.Err)
}

// NewAppError creates a new AppError
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// NotFoundError creates a 404 Not Found error
func NotFoundError(message string) *AppError {
	return NewAppError(http.StatusNotFound, message, nil)
}

// InternalServerError creates a 500 Internal Server Error
func InternalServerError(err error) *AppError {
	return NewAppError(http.StatusInternalServerError, "Internal Server Error", err)
} 
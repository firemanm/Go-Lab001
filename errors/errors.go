package errors

import (
	"fmt"
)

type AppError struct {
	Msg  string
	Code int
}

func NewAppError(msg string, code int) *AppError {
	return &AppError{Msg: msg, Code: code}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

package errors

import "gin-blog/pkg/globals"

type AppError struct {
	Code    globals.AppCode `json:"code"`
	Message string          `json:"error"`
}

func NewAppeError(code globals.AppCode, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

func (e *AppError) Error() string { return e.Message }

package errors

import "fmt"

type ErrorCode int

const (
	ErrInvalidInput ErrorCode = iota
	ErrNotFound
	ErrInternal
)

type AppError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewInvalidInputError(message string) error {
	return &AppError{
		Code:    ErrInvalidInput,
		Message: message,
	}
}

func NewNotFoundError(message string) error {
	return &AppError{
		Code:    ErrNotFound,
		Message: message,
	}
}

func NewInternalError(message string, err error) error {
	return &AppError{
		Code:    ErrInternal,
		Message: message,
		Err:     err,
	}
}

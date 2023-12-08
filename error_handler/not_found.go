package error_handler

import "fmt"

const NotFoundErr = "NotFoundError"

type NotFoundError struct {
	ErrCode     string
	Message     string
	OriginalErr string
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{
		ErrCode: NotFoundErr,
		Message: msg,
	}
}

func NewNotFoundErrorWithOriginalErr(msg string, originalErr string) *NotFoundError {
	return &NotFoundError{
		Message:     msg,
		OriginalErr: originalErr,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", NotFoundErr, e.Message)
}

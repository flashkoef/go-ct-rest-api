package errors

import "fmt"

const NotFoundErr = "NotFoundError"

type NotFoundError struct {
	Message string
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{Message: msg}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", NotFoundErr, e.Message)
}

package error_handler

import "fmt"

const InternalErr = "InternalError"

type InternalError struct {
	Message string
	Err     error
}

func NewInternalError(msg string, err error) *InternalError {
	return &InternalError{
		Message: msg,
		Err:     err,
	}
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%s: Internal error. original error: %s", e.Message, e.Err)
}

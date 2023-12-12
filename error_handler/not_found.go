package error_handler

import "fmt"

const NotFoundErr = "NotFoundError"

type NotFoundError struct {
	ErrCode string `json:"errCode"`
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{
		ErrCode: NotFoundErr,
		Message: msg,
	}
}

func NewNotFoundErrorWithOriginalErr(msg string, originalErr string) *NotFoundError {
	return &NotFoundError{
		Message: msg,
		Err:     originalErr,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", NotFoundErr, e.Message)
}

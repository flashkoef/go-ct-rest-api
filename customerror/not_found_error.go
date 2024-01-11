package customerror

import "fmt"

const NotFoundErr = "NotFoundError"

type NotFoundError struct {
	ErrCode string `json:"errCode"`
	Message string `json:"message"`
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{
		ErrCode: NotFoundErr,
		Message: msg,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", NotFoundErr, e.Message)
}

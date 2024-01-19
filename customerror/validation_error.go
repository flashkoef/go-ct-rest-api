package customerror

import "fmt"

const ValidationErr = "ValidationError"

type ValidationError struct {
	ErrCode string `json:"errCode"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func NewValidationError(msg string, err error) *ValidationError {
	return &ValidationError{
		ErrCode: ValidationErr,
		Message: msg,
		Err:     err,
	}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s, error: %s", e.Message, e.Err)
}

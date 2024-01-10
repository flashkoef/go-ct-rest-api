package errorhandler

import "fmt"

const InternalErr = "InternalError"

type InternalError struct {
	ErrCode string `json:"errCode"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func NewInternalError(msg string, err error) *InternalError {
	return &InternalError{
		ErrCode: InternalErr,
		Message: msg,
		Err:     err,
	}
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%s: Internal error. original error: %s", e.Message, e.Err)
}

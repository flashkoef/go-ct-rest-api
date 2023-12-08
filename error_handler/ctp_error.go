package error_handler

import "fmt"

const CtpErr = "CommercetoolsError"

type CtpError struct {
	ErrCode string
	Message string
	Err     error
}

func NewCtpError(msg string, err error) *CtpError {
	return &CtpError{
		ErrCode: CtpErr,
		Message: msg,
		Err:     err,
	}
}

func (e *CtpError) Error() string {
	return fmt.Sprintf("%s: %s: %s", CtpErr, e.Message, e.Err) // "%s: commercetools platform error. error: %s", e.Message, e.Err
}

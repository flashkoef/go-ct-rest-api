package errorhandler

import "fmt"

const CtpErr = "CommercetoolsError"

type CtpError struct {
	ErrCode string `json:"errCode"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func NewCtpError(msg string, err error) *CtpError {
	return &CtpError{
		ErrCode: CtpErr,
		Message: msg,
		Err:     err,
	}
}

func (e *CtpError) Error() string {
	// "%s: commercetools platform error. error: %s", e.Message, e.Err
	return fmt.Sprintf("%s: %s: %s", CtpErr, e.Message, e.Err)
}

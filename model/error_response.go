package model

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Error   error  `json:"error"`
}

func NewErrorResponse(msg string, code string, err error) *ErrorResponse {
	return &ErrorResponse{
		Message: msg,
		Code:    code,
		Error:   err,
	}
}

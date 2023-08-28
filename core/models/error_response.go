package models

type ErrorResponse struct {
	Status uint `json:"status"`
	ErrorCode string `json:"errorCode"`
	Error error `json:"error"`
}

func NewErrorResponse(status uint, errCode string, err error) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		ErrorCode: errCode,
		Error: err,
	}
}

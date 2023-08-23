package models

type NotFoundRes struct {
	Status uint `json:"status"`
	Err string `json:"err"`
	Msg string `json:"msg"`
}

func NewNotFoundRes(msg string) *NotFoundRes {
	return &NotFoundRes{
		Status: 404,
		Err: "NOT_FOUND_ERROR",
		Msg: msg,
	}
}

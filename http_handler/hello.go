package http_handler

import "github.com/gin-gonic/gin"

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (handler *HelloHandler) Hello(ctx *gin.Context) {
	ctx.JSON(200, struct {
		Message string `json:"message"`
	}{Message: "Hello, from your API!"})
}

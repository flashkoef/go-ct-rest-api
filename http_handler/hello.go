package http_handler

import "github.com/gin-gonic/gin"

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) Hello(ctx *gin.Context) {
	ctx.JSON(200, struct {
		Message string `json:"message"`
	}{Message: "Hello, from your API!"})
}

package v1

import "github.com/gin-gonic/gin"

// type Controller struct {
// 	serviceImpl v1.Servicer
// }

type HelloController struct {}

// func NewHelloController(s *v1.Service) *Controller {
// 	return &Controller{serviceImpl: s}
// }

func NewHelloController() *HelloController {
	return &HelloController{}
}

// func (c *HelloController) Hello(ctx *gin.Context) {
// 	data := c.serviceImpl.HelloService()
// 	ctx.JSON(200, gin.H{
// 		"message": data,
// 	})
// }

func (c *HelloController) Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello",
	})
}

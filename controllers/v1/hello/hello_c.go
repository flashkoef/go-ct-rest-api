package hello

import (
	"github.com/flashkoef/go-ct-rest-api/services/v1"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	serviceImpl v1.Servicer
}

func NewHelloController(s *v1.Service) *Controller {
	return &Controller{serviceImpl: s}
}

func (c *Controller) Hello(ctx *gin.Context) {
	data := c.serviceImpl.HelloService()
	ctx.JSON(200, gin.H{
		"message": data,
	})
}

package v1

import "github.com/gin-gonic/gin"

type Controller struct {
	serviceImpl services.Servicer
}

func NewHelloController(s *services.Service) *Controller {
	return &Controller{serviceImpl: s}
}

func (c *Controller) Hello(ctx *gin.Context) {
	c.serviceImpl.HelloService()
}

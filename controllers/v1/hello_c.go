package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/flashkoef/go-ct-rest-api/services/v1"
)

type Controller struct {
	serviceImpl v1.Servicer
}

func NewHelloController(s *v1.Service) *Controller {
	return &Controller{serviceImpl: s}
}

func (c *Controller) Hello(ctx *gin.Context) {
	c.serviceImpl.HelloService()
}

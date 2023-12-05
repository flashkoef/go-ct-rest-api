package router

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1"
	"github.com/gin-gonic/gin"
)

func SetHelloRoute(router *gin.RouterGroup, c *v1.HelloController) {
	router.GET("/hello", c.Hello)
}

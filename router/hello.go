package router

import (
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/gin-gonic/gin"
)

func SetHelloRoute(router *gin.RouterGroup, c *http_handler.HelloController) {
	router.GET("/hello", c.Hello)
}

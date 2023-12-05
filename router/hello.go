package router

import (
	"github.com/flashkoef/go-ct-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetHelloRoute(router *gin.RouterGroup, c *controllers.HelloController) {
	router.GET("/hello", c.Hello)
}

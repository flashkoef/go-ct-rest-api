package v1

import (	
	"github.com/flashkoef/go-ct-rest-api/controllers/v1/hello"
	"github.com/gin-gonic/gin"
)

func SetHelloRoute(router *gin.RouterGroup, c *v1.Controller) {
	router.GET("/hello", c.Hello)
}

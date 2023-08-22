package v1

import (
	"github.com/gin-gonic/gin"
	
	"github.com/flashkoef/go-ct-rest-api/controllers/v1"
)

func SetHelloRoute(router *gin.RouterGroup, c *v1.Controller) {
	router.GET("/hello", c.Hello)
}

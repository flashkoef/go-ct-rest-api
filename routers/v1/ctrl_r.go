package v1

import "github.com/gin-gonic/gin"

func SetHelloRoute(router *gin.RouterGroup, c *ctrl.Controller) {
	router.GET("/hello", c.HelloController)
}

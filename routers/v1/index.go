package v1

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1/hello"
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup, c *v1.Controller) {
	SetHelloRoute(g, c)
}

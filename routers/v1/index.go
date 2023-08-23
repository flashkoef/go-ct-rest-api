package v1

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1/hello"
	hs "github.com/flashkoef/go-ct-rest-api/services/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {	
	SetHelloRoute(g, hello.NewHelloController(hs.NewHelloService()))
}

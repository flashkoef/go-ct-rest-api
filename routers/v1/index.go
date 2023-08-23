package v1

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1/customers"
	"github.com/flashkoef/go-ct-rest-api/controllers/v1/hello"
	connector "github.com/flashkoef/go-ct-rest-api/libs"
	hs "github.com/flashkoef/go-ct-rest-api/services/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {	
	SetHelloRoute(g, hello.NewHelloController(hs.NewHelloService()))
	SetCustomersRoute(g, customers.New(connector.New().GetProjectClient()))
}

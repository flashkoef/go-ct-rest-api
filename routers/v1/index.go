package v1

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1"
	"github.com/flashkoef/go-ct-rest-api/controllers/v1/hello"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	"github.com/flashkoef/go-ct-rest-api/services"
	hs "github.com/flashkoef/go-ct-rest-api/services/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {
	SetHelloRoute(g, hello.NewHelloController(hs.NewHelloService()))
	SetCustomersRoute(g, customers.New(services.New(connector.New().GetProjectClient())))
}

package v1

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1"
	"github.com/flashkoef/go-ct-rest-api/controllers/v1/hello"
	"github.com/flashkoef/go-ct-rest-api/core/errors"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	"github.com/flashkoef/go-ct-rest-api/services"
	hs "github.com/flashkoef/go-ct-rest-api/services/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	SetHelloRoute(rg, hello.NewHelloController(hs.NewHelloService()))
	
	SetCustomersRoute(rg, customers.New(
			services.New(connector.New().GetProjectClient()),
			errors.New(),
		),
	)
}

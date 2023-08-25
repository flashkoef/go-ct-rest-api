package v1

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1"
	"github.com/flashkoef/go-ct-rest-api/core/errors"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	"github.com/flashkoef/go-ct-rest-api/services"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	SetHelloRoute(rg, v1.NewHelloController())

	SetCustomersRoute(rg, v1.NewCustomersController(
		services.New(connector.New().GetProjectClient()),
		errors.New(),
	))
}

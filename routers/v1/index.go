package v1

import (
	controllers "github.com/flashkoef/go-ct-rest-api/controllers/v1"
	"github.com/flashkoef/go-ct-rest-api/core/error_handler"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	services "github.com/flashkoef/go-ct-rest-api/services/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	SetHelloRoute(rg, controllers.NewHelloController())

	SetCustomersRoute(rg, controllers.NewCustomersController(
		services.NewCustomerService(connector.New().GetProjectClient(), error_handler.New()),
		error_handler.New(),
	))
}

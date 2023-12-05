package router

import (
	"github.com/flashkoef/go-ct-rest-api/controllers"
	"github.com/flashkoef/go-ct-rest-api/core/error_handler"
	"github.com/flashkoef/go-ct-rest-api/core/mapper"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	"github.com/flashkoef/go-ct-rest-api/services"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	SetHelloRoute(rg, controllers.NewHelloController())

	SetCustomerRoute(rg, controllers.NewCustomersController(
		services.NewCustomerService(connector.New().GetProjectClient(), error_handler.New()),
		error_handler.New(),
	))

	SetProductRoute(rg, controllers.NewProductController(
		services.NewProductService(connector.New().GetProjectClient(), error_handler.New()),
		error_handler.New(),
		mapper.NewProductMapper(),
	))
}

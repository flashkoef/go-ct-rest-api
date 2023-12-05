package router

import (
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routerGroup *gin.RouterGroup) {
	SetHelloRoute(routerGroup, http_handler.NewHelloController())

	SetCustomerRoute(routerGroup, http_handler.NewCustomersController(
		service.NewCustomerService(connector.New().GetProjectClient(), error_handler.New()),
		error_handler.New(),
	))

	SetProductRoute(routerGroup, http_handler.NewProductController(
		service.NewProductService(connector.New().GetProjectClient(), error_handler.New()),
		error_handler.New(),
		mapper.NewProductMapper(),
	))
}

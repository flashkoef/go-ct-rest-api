package router

import (
	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routerGroup *gin.RouterGroup) {
	SetCustomerRoute(
		routerGroup, http_handler.NewCustomersHandler(
			service.NewCustomerService(connector.New().GetProjectClient(), error_handler.New()),
			error_handler.New(),
			mapper.NewMapper(),
		),
	)
}

package router

import (
	"github.com/flashkoef/go-ct-rest-api/errorhandler"
	"github.com/flashkoef/go-ct-rest-api/httphandler"
	"github.com/flashkoef/go-ct-rest-api/libs/commercetools/connector"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routerGroup *gin.RouterGroup) {
	SetCustomerRoute(
		routerGroup, httphandler.NewCustomersHandler(
			service.NewCustomerService(connector.New().GetProjectClient()),
			errorhandler.New(),
			mapper.NewMapper(),
		),
	)
}

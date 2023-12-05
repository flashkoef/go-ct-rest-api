package router

import (
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/gin-gonic/gin"
)

func SetCustomerRoute(router *gin.RouterGroup, httpHandler *http_handler.CustomerHandler) {
	router.GET("/customer", httpHandler.GetCustomerByEmail)
	router.GET("/customer/:customerID", httpHandler.GetCustomerByID)
}

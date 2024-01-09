package router

import (
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/gin-gonic/gin"
)

func SetCustomerRoute(router *gin.RouterGroup, httpHandler *http_handler.CustomerHandler) {
	router.GET("/customer/:email", httpHandler.GetCustomerByEmail)
	router.POST("/customer/:email", httpHandler.UpsertCustomer)
	router.DELETE("/customer/:email", httpHandler.DeleteCustomerByEmail)
}

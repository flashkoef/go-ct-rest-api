package router

import (
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/gin-gonic/gin"
)

func SetCustomerRoute(router *gin.RouterGroup, c *http_handler.CustomerController) {
	router.GET("/customer", c.GetCustomerByEmail)
	router.GET("/customer/:customerID", c.GetCustomerByID)
}

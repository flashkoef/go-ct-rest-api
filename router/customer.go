package router

import (
	"github.com/flashkoef/go-ct-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetCustomerRoute(router *gin.RouterGroup, c *controllers.CustomerController) {
	router.GET("/customer", c.GetCustomerByEmail)
	router.GET("/customer/:customerID", c.GetCustomerByID)
}

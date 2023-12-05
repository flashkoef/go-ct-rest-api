package router

import (
	v1 "github.com/flashkoef/go-ct-rest-api/controllers/v1"
	"github.com/gin-gonic/gin"
)

func SetCustomerRoute(router *gin.RouterGroup, c *v1.CustomerController) {
	router.GET("/customer", c.GetCustomerByEmail)
	router.GET("/customer/:customerID", c.GetCustomerByID)
}

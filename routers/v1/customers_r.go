package v1

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1"
	"github.com/gin-gonic/gin"
)

func SetCustomersRoute(router *gin.RouterGroup, c *v1.CustomersController) {
	router.GET("/customer", c.GetCustomerByEmail)
	router.GET("/customer/:customerID", c.GetCustomerByID)
}

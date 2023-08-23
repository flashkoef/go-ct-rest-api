package v1

import (
	"github.com/flashkoef/go-ct-rest-api/controllers/v1/customers"
	"github.com/gin-gonic/gin"
)

func SetCustomersRoute(router *gin.RouterGroup, c *customers.Controller) {
	router.GET("/customer", c.GetCustomerById)
}

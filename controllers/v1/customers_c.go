package v1

import (
	"github.com/flashkoef/go-ct-rest-api/core/errors"
	"github.com/flashkoef/go-ct-rest-api/services"
	"github.com/gin-gonic/gin"
)

type CustomersController struct {
	customerService *services.CustomersService
	checkError      *errors.CheckError
}

func NewCustomersController(s *services.CustomersService, ce *errors.CheckError) *CustomersController {
	return &CustomersController{
		customerService: s,
		checkError:      ce,
	}
}

func (c *CustomersController) GetCustomerByEmail(ctx *gin.Context) {
	customer, err := c.customerService.ExecuteGetCustomerByEmailRequest(ctx)

	shouldReturn := c.checkError.CheckError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(200, customer)
}

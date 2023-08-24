package customers

import (
	"github.com/flashkoef/go-ct-rest-api/core/errors"
	"github.com/flashkoef/go-ct-rest-api/services"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	customerService *services.Service
	checkError *errors.CheckError
}

func New(s *services.Service, ce *errors.CheckError) *Controller {
	return &Controller{
		customerService: s,
		checkError: ce,
	}
}

func (c *Controller) GetCustomerByEmail(ctx *gin.Context) {
	customer, err := c.customerService.ExecuteGetCustomerByEmailRequest(ctx)

	shouldReturn := c.checkError.CheckError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(200, customer)
}

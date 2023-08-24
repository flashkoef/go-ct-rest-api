package customers

import (
	"github.com/flashkoef/go-ct-rest-api/core/errors"
	"github.com/flashkoef/go-ct-rest-api/services"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	customerService *services.Service
}

func New(s *services.Service) *Controller {
	return &Controller{
		customerService: s,
	}
}

func (c *Controller) GetCustomerByEmail(ctx *gin.Context) {
	customer, err := c.customerService.ExecuteGetCustomerByEmailRequest(ctx)

	shouldReturn := errors.CheckError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(200, customer)
}

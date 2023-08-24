package customers

import (
	"log"

	"github.com/flashkoef/go-ct-rest-api/core/errors"
	"github.com/flashkoef/go-ct-rest-api/core/models"
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

	shouldReturn := c.checkError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(200, customer)
}

func (c *Controller) checkError(err error, ctx *gin.Context) bool {
	if err != nil {
		switch e := err.(type) {
		case *errors.NotFoundError:
			ctx.JSON(404, models.NewErrorResponse(404, errors.NotFoundErr, err))
			return true
		case *errors.InternalError:
			ctx.JSON(500, models.NewErrorResponse(500, errors.InternalErr, err))
			return true
		default:
			log.Printf("oops, this was unexpected: %s", e)
			ctx.JSON(500, models.NewErrorResponse(500, errors.InternalErr, err))
			return true
		}
	}
	return false
}

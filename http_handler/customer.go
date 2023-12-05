package http_handler

import (
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService service.CustomerServicer
	checkError      error_handler.ErrorHandler
}

func NewCustomersController(s service.CustomerServicer, ce error_handler.ErrorHandler) *CustomerController {
	return &CustomerController{
		customerService: s,
		checkError:      ce,
	}
}

func (c *CustomerController) GetCustomerByEmail(ctx *gin.Context) {
	customer, err := c.customerService.ExecuteGetCustomerByEmailRequest(ctx)

	shouldReturn := c.checkError.CheckInternError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) GetCustomerByID(ctx *gin.Context) {
	customer, err := c.customerService.ExecuteGetCustomerByIdRequest(ctx)

	shouldReturn := c.checkError.CheckInternError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

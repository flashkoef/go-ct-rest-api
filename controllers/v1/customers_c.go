package v1

import (
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/core/error_handler"
	v1 "github.com/flashkoef/go-ct-rest-api/services/v1"
	"github.com/gin-gonic/gin"
)

type CustomersController struct {
	customerService *v1.CustomersService
	checkError      *error_handler.CheckError
}

func NewCustomersController(s *v1.CustomersService, ce *error_handler.CheckError) *CustomersController {
	return &CustomersController{
		customerService: s,
		checkError:      ce,
	}
}

func (c *CustomersController) GetCustomerByEmail(ctx *gin.Context) {
	customer, err := c.customerService.ExecuteGetCustomerByEmailRequest(ctx)

	shouldReturn := c.checkError.CheckInternError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomersController) GetCustomerByID(ctx *gin.Context) {
	customer, err := c.customerService.ExecuteGetCustomerByIdRequest(ctx)

	shouldReturn := c.checkError.CheckInternError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

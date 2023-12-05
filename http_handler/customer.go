package http_handler

import (
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService service.CustomerServicer
	checkError      error_handler.ErrorHandler
}

func NewCustomersHandler(
	customerService service.CustomerServicer,
	checkError error_handler.ErrorHandler,
) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
		checkError:      checkError,
	}
}

func (handler *CustomerHandler) GetCustomerByEmail(ctx *gin.Context) {
	customer, err := handler.customerService.GetCustomerByEmail(ctx)

	shouldReturn := handler.checkError.CheckInternError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (handler *CustomerHandler) GetCustomerByID(ctx *gin.Context) {
	customer, err := handler.customerService.GetCustomerById(ctx)

	shouldReturn := handler.checkError.CheckInternError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

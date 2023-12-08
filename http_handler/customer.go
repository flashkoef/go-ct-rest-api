package http_handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService service.CustomerServicer
	checkError      error_handler.ErrorHandler
	mapper          mapper.CustomerMapper
}

func NewCustomersHandler(
	customerService service.CustomerServicer,
	checkError error_handler.ErrorHandler,
	mapper mapper.CustomerMapper,
) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
		checkError:      checkError,
		mapper:          mapper,
	}
}

func (handler *CustomerHandler) GetCustomerByEmail(ctx *gin.Context) {
	result, err := handler.customerService.GetCustomerByEmail(ctx)

	if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
		return
	}

	if len(result.Results) == 0 {
		msg := fmt.Sprintf("Can't found customer with email %s", ctx.Param("email"))
		log.Println(msg)
		err := error_handler.NewNotFoundError(msg)
		if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
			return
		}
	}

	customer, _ := handler.mapper.MapCtCustomerToCustomer(result.Results[0])

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

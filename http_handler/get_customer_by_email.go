package http_handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/gin-gonic/gin"
)

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

	customer, err := handler.mapper.MapCtCustomerToCustomer(result.Results[0])
	if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

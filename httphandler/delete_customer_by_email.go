package httphandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/errorhandler"
	"github.com/gin-gonic/gin"
)

func (handler *CustomerHandler) DeleteCustomerByEmail(ctx *gin.Context) {
	customerPagedQueryResponse, err := handler.customerService.GetCustomerByEmail(ctx)
	if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
		return
	}

	if len(customerPagedQueryResponse.Results) == 0 {
		msg := fmt.Sprintf("Can't found customer with email %s", ctx.Param("email"))
		log.Println(msg)
		err := errorhandler.NewNotFoundError(msg)
		
		if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
			return
		}
	}

	ctCustomer, err := handler.customerService.DeleteCustomerByID(
		customerPagedQueryResponse.Results[0].ID,
		customerPagedQueryResponse.Results[0].Version,
		ctx,
	)
	if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
		return
	}

	customer, err := handler.mapper.MapCtCustomerToCustomer(*ctCustomer)
	if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

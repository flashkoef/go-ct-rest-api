package httphandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *CustomerHandler) UpsertCustomer(ctx *gin.Context) {
	customerPagedQueryResponse, err := handler.customerService.GetCustomerByEmail(ctx)

	if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
		return
	}

	if len(customerPagedQueryResponse.Results) == 0 {
		customerSignInResult, err := handler.customerService.CreateCustomer(ctx)
		log.Printf("Created customer in ctp, with ID: %s", customerSignInResult.Customer.ID)
		if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
			return
		}

		customer, err := handler.mapper.MapCtCustomerToCustomer(customerSignInResult.Customer)
		if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
			return
		}

		ctx.JSON(http.StatusCreated, customer)
		return
	}

	ctCustomer, err := handler.customerService.UpdateCustomer(customerPagedQueryResponse.Results[0], ctx)
	if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
		return
	}

	customer, err := handler.mapper.MapCtCustomerToCustomer(*ctCustomer)
	if shouldReturn := handler.checkError.CheckInternError(err, ctx); shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

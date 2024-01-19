package httphandler

import (
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/customerror"
	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (handler *CustomerHandler) UpsertCustomer(ctx *gin.Context) {
	var customerToBeValidated model.Customer

	err := ctx.ShouldBindBodyWith(&customerToBeValidated, binding.JSON)
	if err != nil {
		e := customerror.NewValidationError("validation for provided customer failed", err)
		if shouldReturn := handler.errorHandler.CheckError(e, ctx); shouldReturn {
			return
		}
	}

	customerPagedQueryResponse, err := handler.customerService.GetCustomerByEmail(ctx)

	if shouldReturn := handler.errorHandler.CheckError(err, ctx); shouldReturn {
		return
	}

	if len(customerPagedQueryResponse.Results) == 0 {
		customerSignInResult, err := handler.customerService.CreateCustomer(ctx)
		log.Printf("Created customer in ctp, with ID: %s", customerSignInResult.Customer.ID)

		if shouldReturn := handler.errorHandler.CheckError(err, ctx); shouldReturn {
			return
		}

		customer, err := handler.mapper.MapCtCustomerToCustomer(customerSignInResult.Customer)
		if shouldReturn := handler.errorHandler.CheckError(err, ctx); shouldReturn {
			return
		}

		ctx.JSON(http.StatusCreated, customer)
		return
	}

	ctCustomer, err := handler.customerService.UpdateCustomer(customerPagedQueryResponse.Results[0], ctx)
	if shouldReturn := handler.errorHandler.CheckError(err, ctx); shouldReturn {
		return
	}

	customer, err := handler.mapper.MapCtCustomerToCustomer(*ctCustomer)
	if shouldReturn := handler.errorHandler.CheckError(err, ctx); shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

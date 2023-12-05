package service

import (
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (service *CustomerService) GetCustomerById(ctx *gin.Context) (*platform.Customer, error) {
	customer, err := service.ctClient.Customers().WithId(ctx.Param("customerID")).Get().Execute(ctx)

	shouldReturn, err := service.errorHandler.CheckCtSdkErrorForNonPagedResponse(err, ctx)
	if shouldReturn {
		return &platform.Customer{}, err
	}

	return customer, nil
}

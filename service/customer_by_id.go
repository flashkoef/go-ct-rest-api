package service

import (
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (s *CustomerService) GetCustomerById(ctx *gin.Context) (*platform.Customer, error) {
	customer, err := s.projectClient.Customers().WithId(ctx.Param("customerID")).Get().Execute(ctx)

	shouldReturn, err := s.errorHandler.CheckCtSdkErrorForNonPagedResponse(err, ctx)
	if shouldReturn {
		return &platform.Customer{}, err
	}

	return customer, nil
}
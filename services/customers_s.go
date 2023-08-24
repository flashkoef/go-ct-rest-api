package services

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/core/errors"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type Service struct {
	projectClient *platform.ByProjectKeyRequestBuilder
}

func New(pc *platform.ByProjectKeyRequestBuilder) *Service {
	return &Service{
		projectClient: pc,
	}
}

func (s *Service) ExecuteGetCustomerByEmailRequest(ctx *gin.Context) (platform.Customer, error) {
	customer, err := s.projectClient.Customers().
		Get().
		Where([]string{fmt.Sprintf("email=\"%s\"", ctx.Query("email"))}).
		Execute(ctx)

	if err != nil {
		msg := fmt.Sprintf("error while execute request to ctp %s", err)
		log.Println(msg)
		return platform.Customer{}, errors.NewInternalError(msg, err)
	}

	if len(customer.Results) == 0 {
		msg := fmt.Sprintf("can't found customer with email %s", ctx.Query("email"))
		log.Println(msg)
		return platform.Customer{}, errors.NewNotFoundError(msg)
	}

	return customer.Results[0], nil
}

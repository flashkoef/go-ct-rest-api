package v1

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/core/error_handler"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type CustomersService struct {
	projectClient *platform.ByProjectKeyRequestBuilder
	errorHandler *error_handler.CheckError
}

func NewCustomerService(pc *platform.ByProjectKeyRequestBuilder, eh *error_handler.CheckError) *CustomersService {
	return &CustomersService{
		projectClient: pc,
		errorHandler: eh,
	}
}

func (s *CustomersService) ExecuteGetCustomerByEmailRequest(ctx *gin.Context) (platform.Customer, error) {
	customer, err := s.projectClient.Customers().
		Get().
		Where([]string{fmt.Sprintf("email=\"%s\"", ctx.Query("email"))}).
		Execute(ctx)

	if err != nil {
		msg := fmt.Sprintf("error while execute request to ctp %s", err)
		log.Println(msg)
		return platform.Customer{}, error_handler.NewInternalError(msg, err)
	}

	if len(customer.Results) == 0 {
		msg := fmt.Sprintf("can't found customer with email %s", ctx.Query("email"))
		log.Println(msg)
		return platform.Customer{}, error_handler.NewNotFoundError(msg)
	}

	return customer.Results[0], nil
}

func (s *CustomersService) ExecuteGetCustomerByIdRequest(ctx *gin.Context) (*platform.Customer, error) {
	customer, err := s.projectClient.Customers().WithId(ctx.Param("customerID")).Get().Execute(ctx)

	shouldReturn, err := s.errorHandler.CheckCtSdkError(err, ctx)
	if shouldReturn {
		return &platform.Customer{}, err
	}

	return customer, nil
}

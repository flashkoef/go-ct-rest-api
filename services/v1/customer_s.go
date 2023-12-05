package v1

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/core/error_handler"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type CustomerServicer interface {
	ExecuteGetCustomerByEmailRequest(ctx *gin.Context) (platform.Customer, error)
	ExecuteGetCustomerByIdRequest(ctx *gin.Context) (*platform.Customer, error)
}

type CustomerService struct {
	projectClient *platform.ByProjectKeyRequestBuilder
	errorHandler  error_handler.ErrorHandler
}

func NewCustomerService(pc *platform.ByProjectKeyRequestBuilder, eh error_handler.ErrorHandler) *CustomerService {
	return &CustomerService{
		projectClient: pc,
		errorHandler:  eh,
	}
}

func (s *CustomerService) ExecuteGetCustomerByEmailRequest(ctx *gin.Context) (platform.Customer, error) {
	result, ctSdkErr := s.projectClient.Customers().
		Get().
		Where([]string{fmt.Sprintf("email=\"%s\"", ctx.Query("email"))}).
		Execute(ctx)

	shouldReturn, err := checkError(ctSdkErr, result, ctx)
	if shouldReturn {
		return platform.Customer{}, err
	}

	return result.Results[0], nil
}

func checkError(err error, result *platform.CustomerPagedQueryResponse, ctx *gin.Context) (bool, error) {
	if err != nil {
		msg := fmt.Sprintf("error while execute request to ctp %s", err)
		log.Println(msg)
		return true, error_handler.NewInternalError(msg, err)
	}

	if len(result.Results) == 0 {
		msg := fmt.Sprintf("can't found customer with email %s", ctx.Query("email"))
		log.Println(msg)
		return true, error_handler.NewNotFoundError(msg)
	}
	return false, nil
}

func (s *CustomerService) ExecuteGetCustomerByIdRequest(ctx *gin.Context) (*platform.Customer, error) {
	customer, err := s.projectClient.Customers().WithId(ctx.Param("customerID")).Get().Execute(ctx)

	shouldReturn, err := s.errorHandler.CheckCtSdkErrorForNonPagedResponse(err, ctx)
	if shouldReturn {
		return &platform.Customer{}, err
	}

	return customer, nil
}

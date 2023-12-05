package service

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (s *CustomerService) GetCustomerByEmail(ctx *gin.Context) (platform.Customer, error) {
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

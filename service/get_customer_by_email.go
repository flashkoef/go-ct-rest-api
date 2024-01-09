package service

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (service *CustomerService) GetCustomerByEmail(ctx *gin.Context) (*platform.CustomerPagedQueryResponse, error) {
	result, err := service.ctClient.Customers().
		Get().
		Where([]string{fmt.Sprintf("email=\"%s\"", ctx.Param("email"))}).
		Execute(ctx)

	if err != nil {
		msg := fmt.Sprintln("Error while execute request to commercetools platform.")
		log.Println(msg)
		return &platform.CustomerPagedQueryResponse{}, error_handler.NewCtpError(msg, err)
	}

	return result, nil
}

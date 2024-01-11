package service

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/customerror"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (service *CustomerService) CreateCustomer(ctx *gin.Context) (*platform.CustomerSignInResult, error) {
	var customerDraft platform.CustomerDraft
	if err := ctx.BindJSON(&customerDraft); err != nil {
		msg := fmt.Sprintln("Error while created customer draft.")
		log.Println(msg)
		return &platform.CustomerSignInResult{}, customerror.NewInternalError(msg, err)
	}

	customerSignInResult, err := service.ctClient.Customers().Post(customerDraft).Execute(ctx)
	if err != nil {
		msg := fmt.Sprintln("Error while execute create customer request to commercetools platform.")
		log.Println(msg)
		return &platform.CustomerSignInResult{}, customerror.NewCtpError(msg, err)
	}

	return customerSignInResult, nil
}

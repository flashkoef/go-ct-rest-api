package service

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/customerror"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (service *CustomerService) DeleteCustomerByID(
	customerID string,
	version int,
	ctx *gin.Context,
) (*platform.Customer, error) {
	ctCustomer, err := service.ctClient.Customers().WithId(customerID).Delete().Version(version).Execute(ctx)
	if err != nil {
		msg := fmt.Sprintln("Error while execute delete customer request to commercetools platform.")
		log.Println(msg)
		return &platform.Customer{}, customerror.NewCtpError(msg, err)
	}

	return ctCustomer, nil
}

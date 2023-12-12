package service

import (
	"fmt"
	"log"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (service *CustomerService) UpdateCustomer(ctCustomer platform.Customer, ctx *gin.Context) (*platform.Customer, error) {
	var customer model.Customer
	if err := ctx.BindJSON(&customer); err != nil {
		log.Printf("Error while binding customer: %s", err)
		return &platform.Customer{}, err
	}

	customerUpdate := service.ctCustomerUpdate.CreateCustomerUpdate(ctCustomer, customer)

	result, err := service.ctClient.Customers().WithId(ctCustomer.ID).Post(customerUpdate).Execute(ctx)
	if err != nil {
		msg := fmt.Sprintln("Error while execute update customer request to commercetools platform.")
		log.Println(msg)
		return &platform.Customer{}, error_handler.NewCtpError(msg, err)
	}

	return result, nil
}

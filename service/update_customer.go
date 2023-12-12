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

	customerUpdate := createCustomerUpdate(ctCustomer, customer)

	result, err := service.ctClient.Customers().WithId(ctCustomer.ID).Post(customerUpdate).Execute(ctx)
	if err != nil {
		msg := fmt.Sprintln("Error while execute update customer request to commercetools platform.")
		log.Println(msg)
		return &platform.Customer{}, error_handler.NewCtpError(msg, err)
	}

	return result, nil
}

func createCustomerUpdate(ctCustomer platform.Customer, customer model.Customer) platform.CustomerUpdate {
	return platform.CustomerUpdate{
		Version: ctCustomer.Version,
		Actions: []platform.CustomerUpdateAction{
			createGeneralUpdateAction("changeEmail", "email", customer.Email),
			createGeneralUpdateAction("setFirstName", "firstName", customer.FirstName),
			createGeneralUpdateAction("setLastName", "lastName", customer.LastName),
			createGeneralUpdateAction("setDateOfBirth", "dateOfBirth", customer.DateOfBirth),
			createChangeAddressUpdateAction(*ctCustomer.Addresses[0].ID, customer.Addresses[0]),
		},
	}
}

func createGeneralUpdateAction(action string, field string, value string) map[string]interface{} {
	return map[string]interface{}{
		"action": action,
		field:    value,
	}
}

func createChangeAddressUpdateAction(ctCustomerAddressID string, customerAddress model.Address) map[string]interface{} {
	return map[string]interface{}{
		"action":    "changeAddress",
		"addressId": ctCustomerAddressID,
		"address": map[string]interface{}{
			"streetName": customerAddress.StreetName,
			"city":       customerAddress.City,
			"postalCode": customerAddress.PostalCode,
			"country":    "DE",
		},
	}
}

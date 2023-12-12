package update

import (
	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/labd/commercetools-go-sdk/platform"
)

type CtCustomerUpdater interface {
	CreateCustomerUpdate(ctCustomer platform.Customer, customer model.Customer) platform.CustomerUpdate
}

type CtCustomerUpdate struct{}

func NewCtCustomerUpdate() *CtCustomerUpdate {
	return &CtCustomerUpdate{}
}

func (c *CtCustomerUpdate) CreateCustomerUpdate(ctCustomer platform.Customer, customer model.Customer) platform.CustomerUpdate {
	return platform.CustomerUpdate{
		Version: ctCustomer.Version,
		Actions: createCustomerUpdateActions(ctCustomer, customer),
	}
}

func createCustomerUpdateActions(
	ctCustomer platform.Customer,
	customer model.Customer,
) []platform.CustomerUpdateAction {
	var customerUpdateActions []platform.CustomerUpdateAction

	customerUpdateActions = append(
		customerUpdateActions,
		createGeneralUpdateAction("changeEmail", "email", customer.Email),
		createGeneralUpdateAction("setFirstName", "firstName", customer.FirstName),
		createGeneralUpdateAction("setLastName", "lastName", customer.LastName),
		createGeneralUpdateAction("setDateOfBirth", "dateOfBirth", customer.DateOfBirth),
	)

	for index, ctAddress := range ctCustomer.Addresses {
		customerUpdateActions = append(
			customerUpdateActions,
			createChangeAddressUpdateAction(*ctAddress.ID, customer.Addresses[index]),
		)
	}

	return customerUpdateActions
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
			"country":    customerAddress.County,
		},
	}
}

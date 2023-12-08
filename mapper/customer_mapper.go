package mapper

import (
	"encoding/json"
	"log"

	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/labd/commercetools-go-sdk/platform"
)

type CustomerMapper interface {
	MapCtCustomerToCustomer(ctCustomer platform.Customer) model.Customer
}

func (m *Mapper) MapCtCustomerToCustomer(ctCustomer platform.Customer) model.Customer {
	data, err := ctCustomer.MarshalJSON()
	if err != nil {
		log.Fatalf("Error while marshalling ctCustomer: %s", err)
	}

	var mappedCustomer *model.Customer
	err = json.Unmarshal(data, &mappedCustomer)
	if err != nil {
		log.Fatalf("Error while unmarshal ctCustomer to customer: %s", err)
	}

	return *mappedCustomer
}

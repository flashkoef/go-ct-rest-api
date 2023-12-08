package mapper

import (
	"encoding/json"
	"log"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/labd/commercetools-go-sdk/platform"
)

type CustomerMapper interface {
	MapCtCustomerToCustomer(ctCustomer platform.Customer) (model.Customer, error)
}

func (m *Mapper) MapCtCustomerToCustomer(ctCustomer platform.Customer) (model.Customer, error) {
	data, err := ctCustomer.MarshalJSON()
	if err != nil {
		log.Printf("Error while marshalling ctCustomer: %s", err)
		return model.Customer{}, error_handler.NewInternalError("Error while marshalling commercetools customer.", err)
	}

	var mappedCustomer *model.Customer
	err = json.Unmarshal(data, &mappedCustomer)
	if err != nil {
		log.Printf("Error while unmarshal ctCustomer to customer: %s", err)
		return model.Customer{}, error_handler.NewInternalError("Error while unmarshal commercetools customer to customer.", err)
	}

	return *mappedCustomer, nil
}

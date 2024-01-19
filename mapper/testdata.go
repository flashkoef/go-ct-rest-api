package mapper

import (
	"github.com/flashkoef/go-ct-rest-api/model"
	"github.com/labd/commercetools-go-sdk/platform"
)

var (
	firstName  = "John"
	lastName   = "Doe"
	streetName = "Musterstraße"
	city       = "Jena"
	zipCode    = "07743"

	ctAddresses = []platform.Address{
		{
			StreetName: &streetName,
			City:       &city,
			PostalCode: &zipCode,
		},
	}

	expAddresses = []*model.Address{
		{
			StreetName: streetName,
			City:       city,
			PostalCode: zipCode,
		},
	}
)

var ctCustomer = platform.Customer{
	Email:       "john.doe@test.de",
	FirstName:   &firstName,
	LastName:    &lastName,
	DateOfBirth: &platform.Date{Year: 2023, Month: 12, Day: 06},
	Addresses:   ctAddresses,
}

var expCustomer = model.Customer{
	Email:       "john.doe@test.de",
	FirstName:   "John",
	LastName:    "Doe",
	DateOfBirth: "2023-12-06",
	Addresses:   expAddresses,
}

var ctCustomerWithEmptyAddresses = platform.Customer{
	Email:       "john.doe@test.de",
	FirstName:   &firstName,
	LastName:    &lastName,
	DateOfBirth: &platform.Date{Year: 2023, Month: 12, Day: 06},
	Addresses:   []platform.Address{},
}

var expCustomerWithEmptyAddresses = model.Customer{
	Email:       "john.doe@test.de",
	FirstName:   "John",
	LastName:    "Doe",
	DateOfBirth: "2023-12-06",
	Addresses:   []*model.Address{},
}

var ctCustomerWithEmptyFields = platform.Customer{
	Email:       "john.doe@test.de",
	FirstName:   nil,
	LastName:    nil,
	DateOfBirth: nil,
	Addresses:   ctAddresses,
}

var expCustomerWithEmptyFields = model.Customer{
	Email:       "john.doe@test.de",
	FirstName:   "",
	LastName:    "",
	DateOfBirth: "",
	Addresses:   expAddresses,
}

var ctCustomerWithEmptyAddressFields = platform.Customer{
	Email:       "john.doe@test.de",
	FirstName:   &firstName,
	LastName:    &lastName,
	DateOfBirth: &platform.Date{Year: 2023, Month: 12, Day: 06},
	Addresses:   []platform.Address{},
}

var expCustomerWithEmptyAddressFields = model.Customer{
	Email:       "john.doe@test.de",
	FirstName:   "John",
	LastName:    "Doe",
	DateOfBirth: "2023-12-06",
	Addresses:   []*model.Address{},
}

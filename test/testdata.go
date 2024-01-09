package test

import "github.com/flashkoef/go-ct-rest-api/model"

var (
	expectedCustomer = model.Customer{
		Email:       "jane.doe@test.de",
		Password:    "****wRs=",
		FirstName:   "Jane",
		LastName:    "Doe",
		DateOfBirth: "2019-09-07",
		Addresses: []model.Address{
			{
				StreetName: "Musterstraße 12",
				City:       "Musterstadt",
				PostalCode: "01234",
				County:     "DE",
			},
		},
	}

	expectedErrorResponse = model.ErrorResponse{
		Message: "Resource not found.",
		Code:    "NOT_FOUND",
		Error:   "NotFoundError: Can't found customer with email unknown@test.de",
	}
)

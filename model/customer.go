package model

type Customer struct {
	Email string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
}

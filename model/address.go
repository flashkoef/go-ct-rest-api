package model

type Address struct {
	StreetName string `json:"streetName,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}

package model

type Address struct {
	StreetName string `json:"streetName" binding:"required"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country" binding:"required,iso3166_1_alpha2"`
}

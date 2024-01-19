package model

type Customer struct {
	Email       string     `json:"email" binding:"required"`
	Password    string     `json:"password" binding:"required"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	DateOfBirth string     `json:"dateOfBirth"`
	Addresses   []*Address `json:"addresses" binding:"required,dive"`
}

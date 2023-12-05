package model

type Product struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Sku         string      `json:"sku"`
	Attributes  []Attribute `json:"attributes"`
}

type Attribute struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

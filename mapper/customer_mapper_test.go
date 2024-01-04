package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapCtCustomerToCustomer(t *testing.T) {
	m := NewMapper()
	mappedCustomer, _ := m.MapCtCustomerToCustomer(ctCustomer)
	assert.Equal(t, expCustomer, mappedCustomer)
}

func TestMapCtCustomerToCustomerWithEmptyAddresses(t *testing.T) {
	m := NewMapper()
	mappedCustomer, _ := m.MapCtCustomerToCustomer(ctCustomerWithEmptyAddresses)
	assert.Equal(t, expCustomerWithEmptyAddresses, mappedCustomer)
}

func TestCtCustomerToCustomerWithEmptyCustomerFields(t *testing.T) {
	m := NewMapper()
	mappedCustomer, _ := m.MapCtCustomerToCustomer(ctCustomerWithEmptyFields)
	assert.Equal(t, expCustomerWithEmptyFields, mappedCustomer)
}

func TestMapCtCustomerToCustomerWithEmptyAddressFields(t *testing.T) {
	m := NewMapper()
	mappedCustomer, _ := m.MapCtCustomerToCustomer(ctCustomerWithEmptyAddressFields)
	assert.Equal(t, expCustomerWithEmptyAddressFields, mappedCustomer)
}

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type CustomerServicer interface {
	GetCustomerByEmail(ctx *gin.Context) (*platform.CustomerPagedQueryResponse, error)
	UpdateCustomer(ctCustomer platform.Customer, ctx *gin.Context) (*platform.Customer, error)
	CreateCustomer(ctx *gin.Context) (*platform.CustomerSignInResult, error)
	DeleteCustomerByID(customerID string, version int, ctx *gin.Context) (*platform.Customer, error)
}

type CustomerService struct {
	ctClient *platform.ByProjectKeyRequestBuilder
}

func NewCustomerService(ctClient *platform.ByProjectKeyRequestBuilder) *CustomerService {
	return &CustomerService{
		ctClient: ctClient,
	}
}

package service

import (
	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type CustomerServicer interface {
	GetCustomerByEmail(ctx *gin.Context) (platform.Customer, error)
	GetCustomerById(ctx *gin.Context) (*platform.Customer, error)
}

type CustomerService struct {
	projectClient *platform.ByProjectKeyRequestBuilder
	errorHandler  error_handler.ErrorHandler
}

func NewCustomerService(pc *platform.ByProjectKeyRequestBuilder, eh error_handler.ErrorHandler) *CustomerService {
	return &CustomerService{
		projectClient: pc,
		errorHandler:  eh,
	}
}

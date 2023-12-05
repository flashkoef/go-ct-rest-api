package service

import (
	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type ProductServicer interface {
	GetProductBySlug(ctx *gin.Context) (platform.ProductProjection, error)
}

type ProductService struct {
	projectClient *platform.ByProjectKeyRequestBuilder
	errorHandler  error_handler.ErrorHandler
}

func NewProductService(pc *platform.ByProjectKeyRequestBuilder, eh error_handler.ErrorHandler) *ProductService {
	return &ProductService{
		projectClient: pc,
		errorHandler:  eh,
	}
}

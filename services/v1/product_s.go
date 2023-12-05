package v1

import (
	"fmt"

	"github.com/flashkoef/go-ct-rest-api/core/error_handler"
	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

type ProductServicer interface {
	GetProductBySlug(ctx *gin.Context) (platform.ProductProjection, error)
}

type ProductService struct {
	projectClient *platform.ByProjectKeyRequestBuilder
	errorHandler error_handler.ErrorHandler
}

func NewProductService(pc *platform.ByProjectKeyRequestBuilder, eh error_handler.ErrorHandler) *ProductService {
	return &ProductService{
		projectClient: pc,
		errorHandler: eh,
	}
}

func (s *ProductService) GetProductBySlug(ctx *gin.Context) (platform.ProductProjection, error) {
	result, ctSdkErr := s.projectClient.ProductProjections().
		Get().
		Where([]string{fmt.Sprintf("slug(%s=\"%s\")", ctx.Query("language"), ctx.Query("slug"))}).
		Execute(ctx)
	
	shouldReturn, err := s.errorHandler.CheckCtSdkErrorForPagedResponse(ctSdkErr, result, ctx)
	if shouldReturn {
		return platform.ProductProjection{}, err
	}

	return result.Results[0], nil
}

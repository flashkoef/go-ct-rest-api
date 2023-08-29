package v1

import (
	"fmt"
	"log"

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
	
	shouldReturn, err := checkError(ctSdkErr, result, ctx)
	if shouldReturn {
		return platform.ProductProjection{}, err
	}

	return result.Results[0], nil
}

func checkError(err error, result *platform.ProductProjectionPagedQueryResponse, ctx *gin.Context) (bool, error) {
	if err != nil {
		msg := fmt.Sprintf("error while execute request to ctp %s", err)
		log.Println(msg)
		return true, error_handler.NewInternalError(msg, err)
	}

	if len(result.Results) == 0 {
		msg := fmt.Sprintf("can't found product projection with slug %s", ctx.Query("slug"))
		log.Println(msg)
		return true, error_handler.NewNotFoundError(msg)
	}
	return false, nil
}

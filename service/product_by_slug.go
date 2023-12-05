package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/labd/commercetools-go-sdk/platform"
)

func (service *ProductService) GetProductBySlug(ctx *gin.Context) (platform.ProductProjection, error) {
	result, ctSdkErr := service.ctClient.ProductProjections().
		Get().
		Where([]string{fmt.Sprintf("slug(%s=\"%s\")", ctx.Query("language"), ctx.Query("slug"))}).
		Execute(ctx)

	shouldReturn, err := service.errorHandler.CheckCtSdkErrorForPagedResponse(ctSdkErr, result, ctx)
	if shouldReturn {
		return platform.ProductProjection{}, err
	}

	return result.Results[0], nil
}

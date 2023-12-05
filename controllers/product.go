package controllers

import (
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductServicer
	checkError error_handler.ErrorHandler
	productMapper mapper.IProductMapper
}

func NewProductController(
	s services.ProductServicer, 
	ce error_handler.ErrorHandler, 
	pm mapper.IProductMapper,
) *ProductController {
	return &ProductController{
		productService: s,
		checkError: ce,
		productMapper: pm,
	}
}

func (c *ProductController) GetProductBySlug(ctx *gin.Context) {
	productProjection, err := c.productService.GetProductBySlug(ctx)

	shouldReturn := c.checkError.CheckInternError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, c.productMapper.MapToProduct(productProjection))
}

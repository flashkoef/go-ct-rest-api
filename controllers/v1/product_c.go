package v1

import (
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/core/error_handler"
	"github.com/flashkoef/go-ct-rest-api/core/mapper"
	v1 "github.com/flashkoef/go-ct-rest-api/services/v1"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService v1.ProductServicer
	checkError error_handler.ErrorHandler
	productMapper mapper.IProductMapper
}

func NewProductController(
	s v1.ProductServicer, 
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

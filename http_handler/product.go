package http_handler

import (
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/error_handler"
	"github.com/flashkoef/go-ct-rest-api/mapper"
	"github.com/flashkoef/go-ct-rest-api/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductServicer
	checkError     error_handler.ErrorHandler
	productMapper  mapper.ProductMapperPort
}

func NewProductHandler(
	productService service.ProductServicer,
	errorHandler error_handler.ErrorHandler,
	productMapper mapper.ProductMapperPort,
) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		checkError:     errorHandler,
		productMapper:  productMapper,
	}
}

func (handler *ProductHandler) GetProductBySlug(ctx *gin.Context) {
	productProjection, err := handler.productService.GetProductBySlug(ctx)

	shouldReturn := handler.checkError.CheckInternError(err, ctx)
	if shouldReturn {
		return
	}

	ctx.JSON(http.StatusOK, handler.productMapper.MapToProduct(productProjection))
}

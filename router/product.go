package router

import (
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/gin-gonic/gin"
)

func SetProductRoute(router *gin.RouterGroup, httpHandler *http_handler.ProductHandler) {
	router.GET("/product", httpHandler.GetProductBySlug)
}

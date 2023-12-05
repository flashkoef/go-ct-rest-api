package router

import (
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/gin-gonic/gin"
)

func SetProductRoute(router *gin.RouterGroup, c *http_handler.ProductController) {
	router.GET("/product", c.GetProductBySlug)
}

package router

import (
	v1 "github.com/flashkoef/go-ct-rest-api/controllers/v1"
	"github.com/gin-gonic/gin"
)

func SetProductRoute(router *gin.RouterGroup, c *v1.ProductController) {
	router.GET("/product", c.GetProductBySlug)
}

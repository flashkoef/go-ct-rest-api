package router

import (
	"github.com/flashkoef/go-ct-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetProductRoute(router *gin.RouterGroup, c *controllers.ProductController) {
	router.GET("/product", c.GetProductBySlug)
}

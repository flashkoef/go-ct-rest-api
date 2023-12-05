package router

import (
	"github.com/flashkoef/go-ct-rest-api/http_handler"
	"github.com/gin-gonic/gin"
)

func SetHelloRoute(router *gin.RouterGroup, httpHandler *http_handler.HelloHandler) {
	router.GET("/hello", httpHandler.Hello)
}

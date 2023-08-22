package main

import (
	"log"
	"net/http"

	hc "github.com/flashkoef/go-ct-rest-api/controllers/v1/hello"
	hs "github.com/flashkoef/go-ct-rest-api/services/v1"
	router "github.com/flashkoef/go-ct-rest-api/routers/v1"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.New()
	routerGroup := engine.Group("/v1")
	controller := hc.NewHelloController(hs.NewHelloService())
	router.InitRoutes(routerGroup, controller)
}

func main() {
	log.Println("Server running on port: ", 8080)
	http.ListenAndServe(":8080", engine)
}
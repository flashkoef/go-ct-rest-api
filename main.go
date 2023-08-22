package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/config"
	hc "github.com/flashkoef/go-ct-rest-api/controllers/v1/hello"
	router "github.com/flashkoef/go-ct-rest-api/routers/v1"
	hs "github.com/flashkoef/go-ct-rest-api/services/v1"
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
	conf := config.New()

	log.Println("Server running on port: ", conf.HttpServer.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.HttpServer.Port), engine)
}
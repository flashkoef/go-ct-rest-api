package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/config"
	router "github.com/flashkoef/go-ct-rest-api/routers/v1"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.New()
	routerGroup := engine.Group("/v1")
	router.InitRoutes(routerGroup)
}

func main() {
	conf := config.New()

	log.Println("Server running on port: ", conf.HttpServer.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.HttpServer.Port), engine)
}
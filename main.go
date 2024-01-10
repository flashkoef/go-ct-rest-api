package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashkoef/go-ct-rest-api/config"
	"github.com/flashkoef/go-ct-rest-api/router"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.New()
	routerGroup := engine.Group("")
	router.InitRoutes(routerGroup)
}

func main() {
	conf := config.New()

	log.Println("Server running on port: ", conf.HTTPServer.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.HTTPServer.Port), engine)
}

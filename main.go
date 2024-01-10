package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/flashkoef/go-ct-rest-api/config"
	"github.com/flashkoef/go-ct-rest-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	conf := config.New()
	
	routerGroup := engine.Group("")
	router.InitRoutes(routerGroup)
	
	log.Println("Server running on port: ", conf.HTTPServer.Port)
	
	if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.HTTPServer.Port), engine); err != nil {
		log.Fatalf("error while listen and serve: %s", err)
		os.Exit(1)
	}
}

package main

import (
	"log"
	"net/http"

	c "github.com/flashkoef/go-ct-rest-api/controllers/v1"
	s "github.com/flashkoef/go-ct-rest-api/services/v1"
	r "github.com/flashkoef/go-ct-rest-api/routers/v1"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.New()
	version1 := router.Group("/v1")
	controller := c.NewHelloController(s.NewHelloService())
	r.InitRoutes(version1, controller)
}

func main() {
	log.Println("Server running on port: ", 8080)
	http.ListenAndServe(":8080", router)
}
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.New()
	version1 := router.Group("/v1")
	controller := ctrl.NewController(services.NewService())
	v1.InitRoutes(version1, controller)
}

func main() {
	log.Println("Server running on port: ", 8080)
	http.ListenAndServe(":8080", router)
}
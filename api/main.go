package main

import (
	"api/src/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting the application...")
	setGin()

}

func setGin() {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	router.InitRoutes(&r.RouterGroup)
	http.ListenAndServe(":5000", r)
}

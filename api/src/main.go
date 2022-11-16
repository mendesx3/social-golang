package main

import (
	"api/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting the application...")
	setGin()
	log.Println("application UP")
}

func setGin() {
	r := gin.New()
	api := r.Group("")
	gin.ForceConsoleColor()
	r.Use(gin.Logger(), gin.Recovery())
	router.InitRoutes(api)
	r.Run(":5000")
}

package main

import (
	"api/src/config"
	"api/src/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting the application...")
	//setGin()
	config.InitMySql()
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

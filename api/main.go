package main

import (
	"api/src/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting the application...")
	setGin()
	//config.InitMySql()
	log.Println("application UP")
}

func setGin() {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	router.InitRoutes(&r.RouterGroup)
	r.Run(":5000")
}

package main

import (
	"log"
	// "os"

	
	config "createrestful/configs"
	routes "createrestful/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Connect()

	router := gin.Default()

	routes.Routes(router)
	log.Fatal(router.Run(":3000")) 
}
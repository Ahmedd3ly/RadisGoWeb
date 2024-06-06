package main

import (
	"RadisGoWeb/client/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create a new Gin engine
	router := gin.Default()

	// Register routes defined in routes.go
	routes.RegisterRoutes(router)

	// Serve static files
	router.Static("/static", "./static")

	// Serve HTML templates
	router.LoadHTMLGlob("templates/*")

	// Start the HTTP server
	log.Fatal(router.Run(":8080"))
}

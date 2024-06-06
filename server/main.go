package main

import (
	"RadisGoWeb/server/database"
	"RadisGoWeb/server/middlewares"
	"RadisGoWeb/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()
	server := gin.Default()
	server.Use(middlewares.EnableCORS)

	routes.RegisterRoutes(server)
	server.Run(":8081") // localhost:8080
}

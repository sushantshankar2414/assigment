// main.go
package main

import (
	"event-ticket-reservation/config"
	"event-ticket-reservation/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.ConnectDB()

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
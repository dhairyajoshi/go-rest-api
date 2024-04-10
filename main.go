package main

import (
	"github.com/dhairyajoshi/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8000")
}

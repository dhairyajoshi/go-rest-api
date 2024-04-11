package main

import (
	"github.com/dhairyajoshi/go-rest-api/helpers/database"
	"github.com/dhairyajoshi/go-rest-api/models"
	"github.com/dhairyajoshi/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.GetPostgresConnection()

	server := gin.Default()

	models.MigrateBlogs()

	routes.RegisterRoutes(server)

	server.Run(":8000")
}

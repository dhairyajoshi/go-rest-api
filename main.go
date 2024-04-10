package main

import (
	"github.com/gin-gonic/gin"
	"github.com/dhairyajoshi/routes"
)

func main(){
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8000")
}

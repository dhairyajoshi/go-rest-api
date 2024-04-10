package routes

import (
	"github.com/dhairyajoshi/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(server *gin.Engine) {
	auth := server.Group("/auth/")

	auth.POST("signup/", controllers.UserSignUp)

	auth.POST("login/", controllers.UserLogin)
}

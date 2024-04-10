package routes

import "github.com/gin-gonic/gin"
import "github.com/dhairyajoshi/controllers"

func registerAuthRoutes(server *gin.Engine){
	auth := server.Group("/auth/")

	auth.POST("signup/", controllers.UserSignUp)

	auth.POST("login/", controllers.UserLogin)
}


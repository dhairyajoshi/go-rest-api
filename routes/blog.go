package routes

import "github.com/gin-gonic/gin"
import "github.com/dhairyajoshi/controllers"

func registerBlogRoutes(server *gin.Engine){
	blog := server.Group("/blogs/")

	blog.POST("", controllers.AddBlog)

	blog.GET("", controllers.GetBlogs)
}


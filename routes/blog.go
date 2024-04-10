package routes

import (
	"github.com/dhairyajoshi/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func registerBlogRoutes(server *gin.Engine) {
	blog := server.Group("/blogs/")

	blog.POST("", controllers.AddBlog)

	blog.GET("", controllers.GetBlogs)
}

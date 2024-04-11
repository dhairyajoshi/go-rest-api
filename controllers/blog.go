package controllers

import (
	"strconv"

	"github.com/dhairyajoshi/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func AddBlog(context *gin.Context) {
	var blog models.Blog

	context.ShouldBindJSON(&blog)

	blog.Save()

	context.JSON(200, gin.H{"msg": "Added a blog!"})
}

func GetBlogs(context *gin.Context) {

	blogs := models.Blog{}.Get()

	var response []map[string]any

	for _, blog := range blogs {
		response = append(response, gin.H{"title": blog.Title, "content": blog.Content})
	}

	context.JSON(200, response)
}

func GetBlog(context *gin.Context) {
	id := context.Param("id")

	idx, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.AbortWithStatusJSON(400, gin.H{"msg": "Invalid id!"})
		return
	}

	blogs := models.Blog{}.Get(idx)

	if len(blogs) == 0 {
		context.AbortWithStatusJSON(404, gin.H{"msg": "No blog with given id!"})

		return
	}

	var response map[string]any

	for _, blog := range blogs {
		response = gin.H{"title": blog.Title, "content": blog.Content}
	}

	context.JSON(200, response)
}

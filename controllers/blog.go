package controllers

import "github.com/gin-gonic/gin"

func AddBlog(context *gin.Context) {

	context.JSON(200, gin.H{"msg": "Added a blog!"})
}

func GetBlogs(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "Got all the blogs"})
}
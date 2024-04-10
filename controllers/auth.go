package controllers

import "github.com/gin-gonic/gin"

func UserSignUp(context *gin.Context) {

	context.JSON(200, gin.H{"msg": "sign up route"})
}

func UserLogin(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "login route"})
}
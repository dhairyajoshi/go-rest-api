package controllers

import (
	"github.com/dhairyajoshi/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Email     string
	Password  string
	Cpassword string
}

func UserSignUp(context *gin.Context) {
	User := models.User{}

	var request requestBody

	err := context.ShouldBindJSON(&request)

	if err != nil {
		print(err)
		return
	}

	if request.Password != request.Cpassword {
		context.AbortWithStatusJSON(400, gin.H{"msg": "password and confirm password do not match!"})
		return
	}

	if User.Exists(request.Email) {
		context.AbortWithStatusJSON(400, gin.H{"msg": "User with this email already exists!"})
		return
	}

	var user models.User

	user.Email = request.Email
	user.Password = request.Password

	user.Save()

	context.JSON(200, gin.H{"msg": "User created successfully!"})
}

func UserLogin(context *gin.Context) {
	User := models.User{}

	var request requestBody

	context.ShouldBindJSON(&request)

	_, err := User.Authenticate(request.Email, request.Password)

	if err != nil {
		context.AbortWithStatusJSON(400, gin.H{"msg": "Invalid credentials!"})
		return
	}
	context.JSON(200, gin.H{"msg": "login successful"})
}

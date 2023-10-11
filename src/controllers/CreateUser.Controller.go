package controllers
import (
	//gin 
	"github.com/gin-gonic/gin"
	//package requests
	"onlymachiavelli/web-service-gin/src/requests"
)


func Create (context *gin.Context){
	
	//create a variable to store the request body
	var request requests.CreateUserRequest
	
	//bind the request body to the CreateUserRequest struct
	context.BindJSON(&request)
	
	//return the request body
	context.JSON(200, gin.H{
		"message": request,
	})

}
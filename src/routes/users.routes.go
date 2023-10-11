package routes 


import (

	"github.com/gin-gonic/gin"
	"onlymachiavelli/web-service-gin/src/controllers"
	"fmt"

)


func UsersRoutes(router *gin.Engine){
	
	//to test my server in case it's serverless 
	fmt.Println("Hit the endploit")
	router.GET("/user", controllers.Create)

}

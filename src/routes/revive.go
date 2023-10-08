package routes 


import (
	//gin
	"github.com/gin-gonic/gin"
	//controllers	
	"onlymachiavelli/web-service-gin/src/controllers"
)
func Routes(router *gin.Engine){
	router.GET("/revive", controllers.Revive)
}
// Path: src/controllers/Revive.go
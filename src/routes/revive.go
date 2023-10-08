package routes
import (
	"github.com/gin-gonic/gin"
	"onlymachiavelli/web-service-gin/src/controllers"
)
func Routes(router *gin.Engine){

	//to test my server in case it's serverless 
	router.GET("/revive", controllers.Revive)
}

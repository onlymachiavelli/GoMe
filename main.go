package main

import (
	"onlymachiavelli/web-service-gin/src/helper"
	"onlymachiavelli/web-service-gin/src/utils"
	"onlymachiavelli/web-service-gin/src/routes"
	"github.com/gin-gonic/gin"
	"context"
)

func main() {

	r := gin.Default()
	
	client, err := utils.Connect()
	if err != nil {
		helper.ErrorHandler(err)
	}
	routes.Routes(r)

	
	defer client.Disconnect(context.TODO())

}

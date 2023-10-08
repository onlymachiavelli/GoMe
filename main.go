package main

import (
	//"fmt"
	"onlymachiavelli/web-service-gin/src/helper"

	//"utils"
	"onlymachiavelli/web-service-gin/src/utils"
	"onlymachiavelli/web-service-gin/src/routes"
	//"os"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	"context"
)

func main() {
	//PORT := os.Getenv("PORT")
	// if (!PORT){
	// 	fmt.Println("PORT not found")
	// 	os.Exit(1)
	// }

	//launch the server on the port 
	r := gin.Default()
	//r.Run(":" + PORT)
	
	client, err := utils.Connect()
	if err != nil {
		helper.ErrorHandler(err)
	}
	//middle wares 
	routes.Routes(r)

	
	defer client.Disconnect(context.TODO())

}

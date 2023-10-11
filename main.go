package main

import (
	"onlymachiavelli/web-service-gin/src/helper"
	"onlymachiavelli/web-service-gin/src/utils"
	"onlymachiavelli/web-service-gin/src/routes"
	"github.com/gin-gonic/gin"
	"context"
	"log"
)

func main() {
    r := gin.Default()

    client, err := utils.Connect()
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
        return
    }
	
    defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            log.Printf("Error disconnecting from MongoDB: %v", err)
        }
    }()

    routes.Routes(r)

    // Start the Gin server
    err = r.Run(); // Change the port as needed
    if err != nil {
        //use the helper 
		helper.ErrorHandler(err)
    }
}
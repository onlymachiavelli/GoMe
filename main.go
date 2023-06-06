package main

import (
	"fmt"
	"onlymachiavelli/web-service-gin/helper"
	//"utils"
	"github.com/joho/godotenv"	
	"os"
	"onlymachiavelli/web-service-gin/connector"
	"github.com/gin-gonic/gin"
)


func main () {
	fmt.Println("hello")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}
	PORT:=os.Getenv("PORT") 
	db, err := connector.Connect()
	if (err != nil){
		helper.errorHandler(err)
	}
	defer(db.Close())
	else {
		fmt.Println("Connected to the database")
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	}
	)
	r.Run(":"+PORT)

	
}
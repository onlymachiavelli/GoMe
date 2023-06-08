package main

import (
	"fmt"
	"onlymachiavelli/web-service-gin/src/helper"

	//"utils"
	"onlymachiavelli/web-service-gin/src/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}
	PORT := os.Getenv("PORT")
	db, err := utils.Connect()
	//defer close the db
	fmt.Println(db)
	if err != nil {
		helper.ErrorHandler(err)
	} else {
		fmt.Println("Connected to the database")
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":" + PORT)

}

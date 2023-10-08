package controllers


import (
	//gin 
	"github.com/gin-gonic/gin"

)

//revive controller 
func Revive(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Revive",
	})
}
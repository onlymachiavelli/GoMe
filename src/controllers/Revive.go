package controllers


import (
	"github.com/gin-gonic/gin"
)

func Revive(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Revive",
	})
}
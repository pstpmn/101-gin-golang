package controllers

import (
	"github.com/gin-gonic/gin"
	"example.com/v1/src/repositories"
)

func Index(c *gin.Context) {
	result := repositories.Index();
	c.JSON(200, gin.H{
		"data": result,
		"code":"0",
	})
}
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hello World",
		"code":"0",
	})
}

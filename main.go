package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/john", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"john": "doe",
		})
	})

	r.GET("/john1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"john": "doe",
		})
	})

	r.Run()
}

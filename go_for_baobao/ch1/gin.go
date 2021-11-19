package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "use get method"})
	})
	router.Run(":80")
}

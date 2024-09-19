package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is a simple API",
		})
	})

	// $ go run main.go
	app.Run()
}

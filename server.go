// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Define a route for a GET request
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})
  r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to go programming hello word",
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
}

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	// Simple Health check Endpoint
	r.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Main Endpoint
	r.GET(("/"), func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Microsage!",
		})
	})

	r.Run(":8080")
}


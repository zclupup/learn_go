package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen on 8928 port
	router.Run(":8928") // listens on
	// router.Run() // listens on 0.0.0.0:8080 by default
}

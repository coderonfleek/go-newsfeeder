package main

import "github.com/gin-gonic/gin"

func main() {
	//Create a new 'gin' with some default helpful middleware
	r := gin.Default()

	//Declare a 'GET' route that returns the specified json
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

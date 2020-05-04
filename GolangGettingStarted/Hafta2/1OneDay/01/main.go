package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"firstName": "Kamil",
			"lastName":  "KAPLAN",
			"age":       25,
			"message":   "Merhaba",
		})
	})
	r.Run() // listen and serve on localhost:8080
}

package main

//import "fmt"

import (
	"github.com/gin-gonic/gin"
	"github.com/robbyklein/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()

}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:3000
}

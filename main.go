package main


import (
	"github.com/gin-gonic/gin"
	"github.com/robbyklein/go-crud/controllers"
	"github.com/robbyklein/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()

	initializers.ConnetToDB()

}
func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)       //Create
	r.PUT("/posts", controllers.PostUpdate)        //Update
	r.GET("/posts/:id", controllers.SinglePost)    //Read
	r.GET("/posts", controllers.PostIndex)         //Read multiple
	r.DELETE("/posts/:id", controllers.PostDelete) //Delete

	r.Run()
}

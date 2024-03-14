package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robbyklein/go-crud/initializers"
	"github.com/robbyklein/go-crud/models"
)

func PostCreate(c *gin.Context) {
	//Get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	//Get POST

	var post []models.Post
	initializers.DB.Find(&post)

	//Response

	c.JSON(200, gin.H{
		"post": post,
	})

}

func SinglePost(c *gin.Context) {

	//id from URL
	id := c.Param("id")

	//Get Post
	var post []models.Post
	initializers.DB.First(&post, id)

	//Response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get id
	id := c.Param("id")

	//Get data from req
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//find the data
	var post []models.Post
	initializers.DB.First(&post, id)

	//updation
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Response with update
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {
	//get id
	id := c.Param("id")

	//Deletion
	initializers.DB.Delete(&models.Post{}, id)

	//Response with update
	c.Status(200)

}

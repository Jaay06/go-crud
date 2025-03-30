package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jaay06/go-crud/initializers"
	"github.com/jaay06/go-crud/models"
)

func PostCreate(c *gin.Context) {


	//get request body
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

result := initializers.DB.Create(&post) // pass pointer of data to Create

//eror layer
if result.Error != nil{
	c.Status(400)
	return
}
	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context){

	//get the post
	var posts []models.Post

	results := initializers.DB.Find(&posts)

	if results.Error != nil {
		c.Status(400)
		return
	}

	//return the post index
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context){

	//get id off url
	id := c.Param("id")
	//get the post
	var post models.Post

	results := initializers.DB.First(&post, id)

	if results.Error != nil {
		c.Status(400)
		return
	}

	//return the post index
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context){
	//get the id off the url

	id := c.Param("id")

	//get the data off the req body
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)

	// post := models.Post{Title: body.Title, Body: body.Body}

	//find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)


	//update it
	updateResult := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	if updateResult.Error != nil{
		c.Status(400)
		return
	}

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context){
	//get the id off the url
	id := c.Param("id")
	
	var post models.Post
	//delete the posts
	results := initializers.DB.Delete(&post, id)

	if results.Error != nil {
		c.Status(400)
		return
	}

	//respond
	c.Status(200)
}
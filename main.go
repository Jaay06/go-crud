package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaay06/go-crud/controllers"
	"github.com/jaay06/go-crud/initializers"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main(){
	router := gin.Default()
	
	router.POST("/posts", controllers.PostCreate)
	router.PUT("/posts/:id", controllers.PostsUpdate)

	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)

	router.DELETE("/posts/:id", controllers.PostsDelete)


	
	router.Run() // listen and serve on 0.0.0.0:8080
}
package main

import (
	"github.com/jaay06/go-crud/initializers"
	"github.com/jaay06/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
initializers.DB.AutoMigrate(&models.Post{})
}
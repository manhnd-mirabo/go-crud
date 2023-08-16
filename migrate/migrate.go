package main

import (
	"crud-web/initializers"
	"crud-web/models"
)

func init() {
	initializers.LoadEnvVariales()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
package main

import (
	"github.com/gin-gonic/gin"
	"crud-web/initializers"
	"crud-web/controllers"
)

func init() {
	initializers.LoadEnvVariales()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)

	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
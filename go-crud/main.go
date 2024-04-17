package main

import (
	"github.com/gin-gonic/gin"
	"go-crud/controllers"
	"go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)

	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)

	r.DELETE("/posts/:id", controllers.PostDelete)

	err := r.Run()
	if err != nil {
		return
	}
}

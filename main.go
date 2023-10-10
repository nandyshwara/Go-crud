package main

import (
	"Go-cruud/controllers"
	"Go-cruud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.DbConnect()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}

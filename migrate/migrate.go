package main

import (
	"Go-cruud/initializers"
	"Go-cruud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.DbConnect()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

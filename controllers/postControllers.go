package controllers

import (
	"Go-cruud/initializers"
	"Go-cruud/models"
	"time"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title  string
		Body   string
		Author string
	}
	c.Bind(&body)
	post := models.BlogPost{
		Title:      body.Title,
		Body:       body.Body,
		Author:     body.Author,
		Likes:      0,
		Created_At: time.Now(), // Set the creation timestamp to the current time
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.BlogPost
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	var post models.BlogPost
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title  string
		Body   string
		Author string
	}
	c.Bind(&body)

	var post models.BlogPost
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.BlogPost{
		Title:  body.Title,
		Body:   body.Body,
		Author: body.Author,
	})

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.BlogPost{}, id)

	c.Status(200)
}

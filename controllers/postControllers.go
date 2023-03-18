package controllers

import (
	"net/http"

	"github.com/AkshayAwate/go-crud/initializers"
	"github.com/AkshayAwate/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	var post models.Post
	// Get id from URL
	id := c.Param("id")
	initializers.DB.First(&post, id)
	c.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	var post models.Post

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Get id from URL
	id := c.Param("id")
	initializers.DB.First(&post, id)

	// Update
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)

	c.Status(http.StatusOK)
}

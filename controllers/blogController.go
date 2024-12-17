package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lamadev101/blog-rest-api/initializers"
	"github.com/lamadev101/blog-rest-api/models"
	"github.com/lamadev101/blog-rest-api/utils"
)

func BlogCreate(c *gin.Context) {
	userID := c.GetHeader("UserID")

	// Define a request body struct
	type BlogRequest struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	var req BlogRequest
	c.Bind(&req)

	// if err := c.Bind(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Create the blog object
	blog := models.Blog{
		Title:    req.Title,
		Content:  req.Content,
		Slug:     utils.GenerateSlug(req.Title), // Automatically generate a slug
		AuthorID: userID,
	}

	// Save the blog to the database
	if err := initializers.DB.Create(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"data":    blog,
	})
}

func Blogs(c *gin.Context) {
	var blogs []models.Blog

	if err := initializers.DB.Preload("Author").Find(&blogs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"eror": "Something went wrong"})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": blogs,
	})
}

func Blog(c *gin.Context) {
	slug := c.Param("slug")

	var blog models.Blog
	// Query using the slug column
	if err := initializers.DB.Preload("Author").Where("slug = ?", slug).First(&blog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": blog,
	})
}

func BlogUpdate(c *gin.Context) {
	slug := c.Param("slug")
	var blog models.Blog

	if err := initializers.DB.Where("slug = ?", slug).First(&blog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not found",
		})
		return
	}

	var req map[string]interface{}
	c.Bind(&req)

	if err := initializers.DB.Model(&blog).Updates(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upate blog", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog updated successfully",
	})
}

func BlogDelete(c *gin.Context) {
	slug := c.Param("slug")

	var blog models.Blog
	// Find the blog by slug
	if err := initializers.DB.Where("slug = ?", slug).First(&blog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not found",
		})
		return
	}

	// Delete the blog
	if err := initializers.DB.Delete(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete blog",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
	})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lamadev101/blog-rest-api/initializers"
	"github.com/lamadev101/blog-rest-api/models"
	"gorm.io/gorm"
)

func CategoryCreate(c *gin.Context) {
	var category models.Category

	// Bind the request body to the Category struct
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	// Validate input (e.g., Name should not be empty)
	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Category Name is required",
		})
		return
	}

	// Save the category to the database using GORM
	if err := initializers.DB.Create(&category).Error; err != nil {
		// Return an error response if something goes wrong
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Record not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create category",
			})
		}
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Category created successfully",
	})
}

func Categories(c *gin.Context) {
	var categories []models.Category

	if err := initializers.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Something went wrong!!"})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func CategoryDelete(c *gin.Context) {
	catId := c.Param("id")

	var category models.Category
	if err := initializers.DB.Where("id=?", catId).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found!",
		})
		return
	}

	if err := initializers.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully!!",
	})
}

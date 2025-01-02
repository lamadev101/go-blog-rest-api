package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lamadev101/blog-rest-api/controllers"
	"github.com/lamadev101/blog-rest-api/initializers"
	"github.com/lamadev101/blog-rest-api/middleware"
)

func init() {
	initializers.ConnectToDb()
	initializers.DbSync()
}

func main() {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Category
	r.POST("/category", controllers.CategoryCreate)
	r.GET("/categories", controllers.Categories)
	r.DELETE("/category/:id", controllers.CategoryDelete)

	// Blogs
	r.POST("/blog", controllers.BlogCreate)
	r.GET("/blog/:slug", controllers.Blog)
	r.GET("/blogs", controllers.Blogs)
	r.PUT("/blog/:slug", controllers.BlogUpdate)
	r.DELETE("/blog/:slug", controllers.BlogDelete)

	// Files Upload
	r.POST("/fileUpload", controllers.HandleFileUpload)
	// Protected routes
	r.GET("/welcome", middleware.RequireAuth, controllers.ProtectedRoute)
	// r.GET("/welcome", controllers.ProtectedRoute)

	// protected := r.Group("/protected")
	// protected.Use(controllers.JWTMiddleware())

	r.Run(":8080")
}

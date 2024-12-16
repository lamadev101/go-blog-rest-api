package initializers

import (
	"log"

	"github.com/lamadev101/blog-rest-api/models"
)

// DbSync initializes and syncs the database schema
func DbSync() {
	// Check if DB is initialized
	if DB == nil {
		log.Fatal("DB is not initialized!")
	}

	// AutoMigrate to sync the models with the database
	if err := DB.AutoMigrate(&models.Blog{}, &models.User{}); err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
	} else {
		log.Println("Database schema synced successfully!")
	}
}

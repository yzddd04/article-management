package main

import (
	"log"
	"os"

	"article-crud-api/config"
	"article-crud-api/handlers"
	"article-crud-api/middleware"
	"article-crud-api/models"
	"article-crud-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configuration")
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Auto migrate database
	if err := models.MigrateDB(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.Default()

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// Initialize handlers
	articleHandler := handlers.NewArticleHandler(db)

	// Setup routes
	routes.SetupRoutes(router, articleHandler)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}

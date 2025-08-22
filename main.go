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
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan konfigurasi default")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}
	defer db.Close()

	if err := models.MigrateDB(db); err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	articleHandler := handlers.NewArticleHandler(db)

	routes.SetupRoutes(router, articleHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di port %s", port)
	log.Fatal(router.Run(":" + port))
}

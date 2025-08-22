package routes

import (
	"article-crud-api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine, articleHandler *handlers.ArticleHandler) {
	// API version 1
	v1 := router.Group("/api/v1")
	{
		// Health check endpoint
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "Article CRUD API is running",
				"version": "1.0.0",
			})
		})

		// Articles endpoints
		articles := v1.Group("/articles")
		{
			articles.GET("", articleHandler.GetAllArticles)    // GET /api/v1/articles
			articles.POST("", articleHandler.CreateArticle)    // POST /api/v1/articles
			articles.GET("/:id", articleHandler.GetArticle)    // GET /api/v1/articles/:id
			articles.PUT("/:id", articleHandler.UpdateArticle) // PUT /api/v1/articles/:id
			articles.DELETE("/:id", articleHandler.DeleteArticle) // DELETE /api/v1/articles/:id
		}
	}

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Article CRUD API",
			"endpoints": gin.H{
				"health": "/api/v1/health",
				"articles": gin.H{
					"get_all":    "GET /api/v1/articles",
					"create":     "POST /api/v1/articles",
					"get_by_id":  "GET /api/v1/articles/:id",
					"update":     "PUT /api/v1/articles/:id",
					"delete":     "DELETE /api/v1/articles/:id",
				},
			},
		})
	})
}

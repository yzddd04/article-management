package routes

import (
	"article-crud-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, articleHandler *handlers.ArticleHandler) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "API Artikel CRUD berjalan",
				"version": "1.0.0",
			})
		})

		articles := v1.Group("/articles")
		{
			articles.GET("", articleHandler.GetAllArticles)
			articles.POST("", articleHandler.CreateArticle)
			articles.GET("/:id", articleHandler.GetArticle)
			articles.PUT("/:id", articleHandler.UpdateArticle)
			articles.DELETE("/:id", articleHandler.DeleteArticle)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Selamat datang di API Artikel CRUD",
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

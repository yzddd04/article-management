package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"article-crud-api/models"

	"github.com/gin-gonic/gin"
)

// ArticleHandler handles HTTP requests for articles
type ArticleHandler struct {
	repo *models.ArticleRepository
}

// NewArticleHandler creates a new article handler
func NewArticleHandler(db *sql.DB) *ArticleHandler {
	return &ArticleHandler{
		repo: models.NewArticleRepository(db),
	}
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// CreateArticle handles POST /articles
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var article models.Article
	
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request data: " + err.Error(),
		})
		return
	}

	// Set current time if not provided
	if article.PublishedAt.IsZero() {
		article.PublishedAt = time.Now()
	}

	if err := h.repo.Create(&article); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to create article: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Article created successfully",
		Data:    article,
	})
}

// GetArticle handles GET /articles/:id
func (h *ArticleHandler) GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid article ID",
		})
		return
	}

	article, err := h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Error:   "Article not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to retrieve article: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Article retrieved successfully",
		Data:    article,
	})
}

// GetAllArticles handles GET /articles
func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to retrieve articles: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Articles retrieved successfully",
		Data:    articles,
	})
}

// UpdateArticle handles PUT /articles/:id
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid article ID",
		})
		return
	}

	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request data: " + err.Error(),
		})
		return
	}

	article.ID = id

	// Check if article exists
	existingArticle, err := h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Error:   "Article not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to check article existence: " + err.Error(),
		})
		return
	}

	// Preserve original timestamps if not provided
	if article.PublishedAt.IsZero() {
		article.PublishedAt = existingArticle.PublishedAt
	}

	if err := h.repo.Update(&article); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to update article: " + err.Error(),
		})
		return
	}

	// Get updated article
	updatedArticle, _ := h.repo.GetByID(id)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Article updated successfully",
		Data:    updatedArticle,
	})
}

// DeleteArticle handles DELETE /articles/:id
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid article ID",
		})
		return
	}

	// Check if article exists
	_, err = h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Error:   "Article not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to check article existence: " + err.Error(),
		})
		return
	}

	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Failed to delete article: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Article deleted successfully",
	})
}

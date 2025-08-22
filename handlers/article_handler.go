package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"article-crud-api/models"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	repo *models.ArticleRepository
}

func NewArticleHandler(db *sql.DB) *ArticleHandler {
	return &ArticleHandler{
		repo: models.NewArticleRepository(db),
	}
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var requestBody interface{}
	
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Data request tidak valid: " + err.Error(),
		})
		return
	}

	// Cek apakah request berupa array atau single object
	switch v := requestBody.(type) {
	case []interface{}:
		// Bulk create untuk array
		h.createMultipleArticles(c, v)
	case map[string]interface{}:
		// Single create untuk object
		h.createSingleArticle(c, v)
	default:
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Format data tidak valid. Gunakan object untuk satu artikel atau array untuk beberapa artikel",
		})
	}
}

func (h *ArticleHandler) createSingleArticle(c *gin.Context, data map[string]interface{}) {
	article := h.mapToArticle(data)
	
	if article.PublishedAt.IsZero() {
		article.PublishedAt = time.Now()
	}

	if err := h.repo.Create(&article); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Gagal membuat artikel: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Artikel berhasil dibuat",
		Data:    article,
	})
}

func (h *ArticleHandler) createMultipleArticles(c *gin.Context, data []interface{}) {
	var createdArticles []models.Article
	var errors []string

	for i, item := range data {
		if articleData, ok := item.(map[string]interface{}); ok {
			article := h.mapToArticle(articleData)
			
			if article.PublishedAt.IsZero() {
				article.PublishedAt = time.Now()
			}

			if err := h.repo.Create(&article); err != nil {
				errors = append(errors, fmt.Sprintf("Artikel ke-%d: %v", i+1, err))
			} else {
				createdArticles = append(createdArticles, article)
			}
		} else {
			errors = append(errors, fmt.Sprintf("Artikel ke-%d: format data tidak valid", i+1))
		}
	}

	if len(errors) > 0 && len(createdArticles) == 0 {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Gagal membuat semua artikel: " + strings.Join(errors, "; "),
		})
		return
	}

	message := fmt.Sprintf("Berhasil membuat %d artikel", len(createdArticles))
	if len(errors) > 0 {
		message += fmt.Sprintf(", %d gagal", len(errors))
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data: map[string]interface{}{
			"created_articles": createdArticles,
			"total_created":    len(createdArticles),
			"total_failed":     len(errors),
			"errors":           errors,
		},
	})
}

func (h *ArticleHandler) mapToArticle(data map[string]interface{}) models.Article {
	article := models.Article{}
	
	if title, ok := data["title"].(string); ok {
		article.Title = title
	}
	if content, ok := data["content"].(string); ok {
		article.Content = content
	}
	if author, ok := data["author"].(string); ok {
		article.Author = author
	}
	if category, ok := data["category"].(string); ok {
		article.Category = category
	}
	if publishedAt, ok := data["published_at"].(string); ok {
		if t, err := time.Parse(time.RFC3339, publishedAt); err == nil {
			article.PublishedAt = t
		}
	}
	
	return article
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "ID artikel tidak valid",
		})
		return
	}

	article, err := h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Error:   "Artikel tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Gagal mengambil artikel: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Artikel berhasil diambil",
		Data:    article,
	})
}

func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Gagal mengambil artikel: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Artikel berhasil diambil",
		Data:    articles,
	})
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "ID artikel tidak valid",
		})
		return
	}

	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Data request tidak valid: " + err.Error(),
		})
		return
	}

	article.ID = id

	existingArticle, err := h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Error:   "Artikel tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Gagal memeriksa artikel: " + err.Error(),
		})
		return
	}

	if article.PublishedAt.IsZero() {
		article.PublishedAt = existingArticle.PublishedAt
	}

	if err := h.repo.Update(&article); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Gagal mengupdate artikel: " + err.Error(),
		})
		return
	}

	updatedArticle, _ := h.repo.GetByID(id)

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Artikel berhasil diupdate",
		Data:    updatedArticle,
	})
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "ID artikel tidak valid",
		})
		return
	}

	_, err = h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Error:   "Artikel tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Gagal memeriksa artikel: " + err.Error(),
		})
		return
	}

	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Gagal menghapus artikel: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Artikel berhasil dihapus",
	})
}

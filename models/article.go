package models

import (
	"database/sql"
	"time"
)

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Category    string    `json:"category"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func MigrateDB(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		category TEXT,
		published_at DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	return err
}

func (r *ArticleRepository) Create(article *Article) error {
	query := `
	INSERT INTO articles (title, content, author, category, published_at, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	result, err := r.db.Exec(query, article.Title, article.Content, article.Author, article.Category, article.PublishedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	article.ID = int(id)
	return nil
}

func (r *ArticleRepository) GetByID(id int) (*Article, error) {
	query := `SELECT id, title, content, author, category, published_at, created_at, updated_at FROM articles WHERE id = ?`

	article := &Article{}
	err := r.db.QueryRow(query, id).Scan(
		&article.ID, &article.Title, &article.Content, &article.Author,
		&article.Category, &article.PublishedAt, &article.CreatedAt, &article.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *ArticleRepository) GetAll() ([]*Article, error) {
	query := `SELECT id, title, content, author, category, published_at, created_at, updated_at FROM articles ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*Article
	for rows.Next() {
		article := &Article{}
		err := rows.Scan(
			&article.ID, &article.Title, &article.Content, &article.Author,
			&article.Category, &article.PublishedAt, &article.CreatedAt, &article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *ArticleRepository) Update(article *Article) error {
	query := `
	UPDATE articles 
	SET title = ?, content = ?, author = ?, category = ?, published_at = ?, updated_at = CURRENT_TIMESTAMP
	WHERE id = ?`

	_, err := r.db.Exec(query, article.Title, article.Content, article.Author, article.Category, article.PublishedAt, article.ID)
	return err
}

func (r *ArticleRepository) Delete(id int) error {
	query := `DELETE FROM articles WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

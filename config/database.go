package config

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./articles.db"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database berhasil terkoneksi")
	return db, nil
}

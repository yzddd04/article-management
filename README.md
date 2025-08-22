# API Artikel CRUD

API RESTful modern untuk mengelola artikel yang dibangun dengan Go, framework Gin, dan database SQLite.

## Fitur

- ✅ Operasi CRUD lengkap untuk artikel
- ✅ Desain API RESTful
- ✅ Database SQLite dengan migrasi otomatis
- ✅ Validasi input dan penanganan error
- ✅ Dukungan CORS
- ✅ Logging request
- ✅ Struktur project modern
- ✅ Siap untuk testing di Postman

## Project Structure

```
article-crud-api/
├── main.go                 # Application entry point
├── go.mod                  # Go module file
├── env.example             # Environment variables example
├── README.md              # This file
├── config/
│   └── database.go        # Database configuration
├── models/
│   └── article.go         # Article model and repository
├── handlers/
│   └── article_handler.go # HTTP request handlers
├── middleware/
│   └── middleware.go      # CORS and logging middleware
└── routes/
    └── routes.go          # API routes configuration
```

## Cara Cepat Memulai

### Persyaratan

- Go 1.21 atau lebih tinggi
- Git

### Instalasi

1. Clone repository:
```bash
git clone <repository-url>
cd article-crud-api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Buat file environment:
```bash
cp env.example .env
```

4. Jalankan aplikasi:
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## API Endpoints

### Base URL
```
http://localhost:8080
```

### Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/` | Pesan selamat datang dan info API |
| GET | `/api/v1/health` | Pengecekan status |
| GET | `/api/v1/articles` | Ambil semua artikel |
| POST | `/api/v1/articles` | Buat artikel baru |
| GET | `/api/v1/articles/:id` | Ambil artikel berdasarkan ID |
| PUT | `/api/v1/articles/:id` | Update artikel |
| DELETE | `/api/v1/articles/:id` | Hapus artikel |

## Dokumentasi API

### Model Artikel

```json
{
  "id": 1,
  "title": "Judul Artikel",
  "content": "Konten artikel...",
  "author": "Nama Penulis",
  "category": "Teknologi",
  "published_at": "2024-01-01T00:00:00Z",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Format Response

Semua response API mengikuti format ini:

```json
{
  "success": true,
  "message": "Operasi berhasil diselesaikan",
  "data": { ... },
  "error": null
}
```

### Response Error

```json
{
  "success": false,
  "message": null,
  "data": null,
  "error": "Deskripsi error"
}
```

## Testing dengan Postman

### 1. Buat Artikel (POST)

**URL:** `http://localhost:8080/api/v1/articles`

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "title": "Memulai dengan Go",
  "content": "Go adalah bahasa pemrograman yang powerful...",
  "author": "John Doe",
  "category": "Pemrograman",
  "published_at": "2024-01-01T00:00:00Z"
}
```

### 2. Ambil Semua Artikel (GET)

**URL:** `http://localhost:8080/api/v1/articles`

### 3. Ambil Artikel berdasarkan ID (GET)

**URL:** `http://localhost:8080/api/v1/articles/1`

### 4. Update Artikel (PUT)

**URL:** `http://localhost:8080/api/v1/articles/1`

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "title": "Judul yang Diupdate",
  "content": "Konten yang diupdate...",
  "author": "Jane Smith",
  "category": "Teknologi",
  "published_at": "2024-01-01T00:00:00Z"
}
```

### 5. Hapus Artikel (DELETE)

**URL:** `http://localhost:8080/api/v1/articles/1`

### 6. Pengecekan Status (GET)

**URL:** `http://localhost:8080/api/v1/health`

## Collection Postman

Anda dapat mengimpor collection ini ke Postman untuk testing yang mudah:

```json
{
  "info": {
    "name": "API Artikel CRUD",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Pengecekan Status",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "http://localhost:8080/api/v1/health",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "v1", "health"]
        }
      }
    },
    {
      "name": "Ambil Semua Artikel",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "http://localhost:8080/api/v1/articles",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "v1", "articles"]
        }
      }
    },
    {
      "name": "Buat Artikel",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"title\": \"Memulai dengan Go\",\n  \"content\": \"Go adalah bahasa pemrograman yang powerful...\",\n  \"author\": \"John Doe\",\n  \"category\": \"Pemrograman\",\n  \"published_at\": \"2024-01-01T00:00:00Z\"\n}"
        },
        "url": {
          "raw": "http://localhost:8080/api/v1/articles",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "v1", "articles"]
        }
      }
    },
    {
      "name": "Ambil Artikel berdasarkan ID",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "http://localhost:8080/api/v1/articles/1",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "v1", "articles", "1"]
        }
      }
    },
    {
      "name": "Update Artikel",
      "request": {
        "method": "PUT",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"title\": \"Judul yang Diupdate\",\n  \"content\": \"Konten yang diupdate...\",\n  \"author\": \"Jane Smith\",\n  \"category\": \"Teknologi\",\n  \"published_at\": \"2024-01-01T00:00:00Z\"\n}"
        },
        "url": {
          "raw": "http://localhost:8080/api/v1/articles/1",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "v1", "articles", "1"]
        }
      }
    },
    {
      "name": "Hapus Artikel",
      "request": {
        "method": "DELETE",
        "header": [],
        "url": {
          "raw": "http://localhost:8080/api/v1/articles/1",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "v1", "articles", "1"]
        }
      }
    }
  ]
}
```

## Variabel Environment

| Variable | Default | Deskripsi |
|----------|---------|-----------|
| `PORT` | `8080` | Port server |
| `GIN_MODE` | `debug` | Mode Gin (debug/release) |
| `DB_PATH` | `./articles.db` | Path database SQLite |

## Pengembangan

### Menjalankan dalam Mode Development

```bash
go run main.go
```

### Build untuk Production

```bash
go build -o article-api main.go
```

### Menjalankan Test

```bash
go test ./...
```

## Database

Aplikasi menggunakan SQLite untuk kemudahan. File database (`articles.db`) akan dibuat otomatis saat pertama kali menjalankan aplikasi.

### Schema Database

```sql
CREATE TABLE articles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    author TEXT NOT NULL,
    category TEXT,
    published_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Kontribusi

1. Fork repository
2. Buat feature branch
3. Lakukan perubahan
4. Tambahkan test jika diperlukan
5. Submit pull request

## Lisensi

Project ini open source dan tersedia di bawah [MIT License](LICENSE).

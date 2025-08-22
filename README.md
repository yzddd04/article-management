# Article CRUD API

A modern, well-structured RESTful API for managing articles built with Go, Gin framework, and SQLite database.

## Features

- ✅ Complete CRUD operations for articles
- ✅ RESTful API design
- ✅ SQLite database with automatic migration
- ✅ Input validation and error handling
- ✅ CORS support
- ✅ Request logging
- ✅ Modern project structure
- ✅ Ready for Postman testing

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

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd article-crud-api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create environment file:
```bash
cp env.example .env
```

4. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Base URL
```
http://localhost:8080
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Welcome message and API info |
| GET | `/api/v1/health` | Health check |
| GET | `/api/v1/articles` | Get all articles |
| POST | `/api/v1/articles` | Create new article |
| GET | `/api/v1/articles/:id` | Get article by ID |
| PUT | `/api/v1/articles/:id` | Update article |
| DELETE | `/api/v1/articles/:id` | Delete article |

## API Documentation

### Article Model

```json
{
  "id": 1,
  "title": "Article Title",
  "content": "Article content...",
  "author": "Author Name",
  "category": "Technology",
  "published_at": "2024-01-01T00:00:00Z",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Response Format

All API responses follow this format:

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { ... },
  "error": null
}
```

### Error Response

```json
{
  "success": false,
  "message": null,
  "data": null,
  "error": "Error description"
}
```

## Testing with Postman

### 1. Create Article (POST)

**URL:** `http://localhost:8080/api/v1/articles`

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "title": "Getting Started with Go",
  "content": "Go is a powerful programming language...",
  "author": "John Doe",
  "category": "Programming",
  "published_at": "2024-01-01T00:00:00Z"
}
```

### 2. Get All Articles (GET)

**URL:** `http://localhost:8080/api/v1/articles`

### 3. Get Article by ID (GET)

**URL:** `http://localhost:8080/api/v1/articles/1`

### 4. Update Article (PUT)

**URL:** `http://localhost:8080/api/v1/articles/1`

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "title": "Updated Title",
  "content": "Updated content...",
  "author": "Jane Smith",
  "category": "Technology",
  "published_at": "2024-01-01T00:00:00Z"
}
```

### 5. Delete Article (DELETE)

**URL:** `http://localhost:8080/api/v1/articles/1`

### 6. Health Check (GET)

**URL:** `http://localhost:8080/api/v1/health`

## Postman Collection

You can import this collection into Postman for easy testing:

```json
{
  "info": {
    "name": "Article CRUD API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Health Check",
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
      "name": "Get All Articles",
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
      "name": "Create Article",
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
          "raw": "{\n  \"title\": \"Getting Started with Go\",\n  \"content\": \"Go is a powerful programming language...\",\n  \"author\": \"John Doe\",\n  \"category\": \"Programming\",\n  \"published_at\": \"2024-01-01T00:00:00Z\"\n}"
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
      "name": "Get Article by ID",
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
      "name": "Update Article",
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
          "raw": "{\n  \"title\": \"Updated Title\",\n  \"content\": \"Updated content...\",\n  \"author\": \"Jane Smith\",\n  \"category\": \"Technology\",\n  \"published_at\": \"2024-01-01T00:00:00Z\"\n}"
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
      "name": "Delete Article",
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

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `GIN_MODE` | `debug` | Gin mode (debug/release) |
| `DB_PATH` | `./articles.db` | SQLite database path |

## Development

### Running in Development Mode

```bash
go run main.go
```

### Building for Production

```bash
go build -o article-api main.go
```

### Running Tests

```bash
go test ./...
```

## Database

The application uses SQLite for simplicity. The database file (`articles.db`) will be created automatically when you first run the application.

### Database Schema

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

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

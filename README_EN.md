## 🇷🇺  [Русская версия](./README.md)

# Log-Stash-Lite

A lightweight and fast service for collecting, storing, and searching logs via REST API. Supports Elasticsearch, JWT
authorization, flexible log structure, and fast full-text search.

**Tech stack:** Go, chi, Elasticsearch, zap, JWT, Docker

---

## 🚀 Quick Start

1. Clone the repository:
   ```sh
   git clone https://github.com/trottling/Log-Stash-Lite.git
   cd Log-Stash-Lite
   ```
2. Start dependencies (Elasticsearch):
   ```sh
   docker-compose up -d elasticsearch
   ```
3. Build and run the service:
   ```sh
   go build -o app ./cmd/main.go
   ./app
   ```
   or via Docker Compose:
   ```sh
   docker-compose up --build
   ```

---

## ⚡️ Key Features

- Accept logs via REST API (POST /add_log, /add_logs)
- Search and filter logs (GET /get_logs, /get_logs_count)
- JWT authorization for protected endpoints
- Flexible log entry structure (JSON)
- Fast full-text and pattern search (Elasticsearch)
- Swagger documentation
- Easy Docker deployment
- Healthcheck and system stats

**Highlights:**

- Minimal dependencies, fast start
- Easily extendable parsers and storage backends
- Support for various log formats (zap, logrus, pino, ...)

---

## 🛠️ Usage / API Example

### Example request (POST /add_log)

```json
{
  "parse_type": "default",
  "log": {
    "level": "info",
    "msg": "hello"
  }
}
```

### Example response

```json
{
  "status": "ok"
}
```

### Example search (GET /get_logs)

```
GET /get_logs?level=info&limit=10
```

### Example response

```json
{
  "logs": [
    {
      "level": "info",
      "msg": "hello"
    }
  ],
  "count": 1
}
```

---

## 🔗 Documentation

- Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Curl examples:
  ```sh
  curl -X POST http://localhost:8080/add_log \
    -H "Authorization: Bearer <token>" \
    -H "Content-Type: application/json" \
    -d '{"parse_type":"default","log":{"msg":"test"}}'
  ```
  ```sh
  curl http://localhost:8080/get_logs?level=info&limit=5 -H "Authorization: Bearer <token>"
  ```

---

## ⚙️ Configuration

| Flag / Variable  | Default               | Description            |
|------------------|-----------------------|------------------------|
| LISTEN_ADDR      | :8080                 | API listen address     |
| ELASTIC_URL      | http://localhost:9200 | Elasticsearch URL      |
| ELASTIC_USERNAME | elastic               | Elasticsearch username |
| ELASTIC_PASSWORD | change_me             | Elasticsearch password |
| LOG_LEVEL        | info                  | Logging level          |
| JWT_SECRET       | changeme              | JWT secret             |

---

## 🐳 Docker / Compose

1. Build the image:
   ```sh
   docker build -t log-stash-lite .
   ```
2. Run with docker-compose:
   ```sh
   docker-compose up --build
   ```
3. Pass parameters:
    - Via environment variables in docker-compose.yml
    - You can mount volumes for logs/config

---

## 🧪 Testing

- Run all tests:
  ```sh
  go test ./...
  ```
- Check coverage:
  ```sh
  go test -cover ./...
  ```
- Integration tests:
  ```sh
  go test -tags=integration ./internal/api/integration
  ```

---

## 📝 FAQ / Examples

- **401 error:** Check your JWT token, use /auth/token to generate one.
- **Elasticsearch not responding:** Check docker-compose and environment variables.
- **How to add a new parser?** See below.
- **How to change storage backend?** Implement the storage.Storage interface and pass it to Handler.

---

## ➕ How to add your own log parser

1. Create a new file in `internal/parser/`, e.g. `mylogger.go`:
   ```go
   package parser

   func ParseMyLogger(log map[string]interface{}) (map[string]interface{}, error) {
       // Your parsing logic here
       return log, nil
   }
   ```
2. Register your parser in `parser.go`:
   ```go
   // ...
   case "mylogger":
       parseFunc = ParseMyLogger
   // ...
   ```
3. Use `"parse_type": "mylogger"` in your API requests.

---

## 💬 Contributing

- Fork the repo, create a branch, submit a PR
- Describe your changes, add tests
- For bugs and features — open an Issue

---

## 📄 License

MIT License. Free to use with attribution.

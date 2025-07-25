## 🇬🇧  [English version](README_EN.md)

# Log-Stash-Lite

Лёгкий и быстрый сервис для сбора, хранения и поиска логов с REST API. Поддерживает Elasticsearch, JWT-авторизацию,
гибкую структуру логов и быстрый full-text поиск.

**Технологии:** Go, chi, Elasticsearch, zap, JWT, Docker

---

## 🚀 Быстрый старт

1. Клонируйте репозиторий:
   ```sh
   git clone https://github.com/trottling/Log-Stash-Lite.git
   cd Log-Stash-Lite
   ```
2. Запустите зависимости (Elasticsearch):
   ```sh
   docker-compose up -d elasticsearch
   ```
3. Соберите и запустите сервис:
   ```sh
   go build -o app ./cmd/main.go
   ./app
   ```
   или через Docker Compose:
   ```sh
   docker-compose up --build
   ```

---

## ⚡️ Основные возможности

- Приём логов через REST API (POST /add_log, /add_logs)
- Поиск и фильтрация логов (GET /get_logs, /get_logs_count)
- JWT-авторизация для защищённых ручек
- Гибкая структура log entry (JSON)
- Быстрый full-text и паттерн-поиск (Elasticsearch)
- Swagger-документация
- Лёгкий запуск через Docker
- Healthcheck и системная статистика

**Особенности:**

- Минимум зависимостей, быстрый старт
- Легко расширять парсеры и хранилища
- Поддержка разных форматов логов (zap, logrus, pino, ...)

---

## 🛠️ Использование / Пример API

### Пример запроса (POST /add_log)

```json
{
  "parse_type": "default",
  "log": {
    "level": "info",
    "msg": "hello"
  }
}
```

### Пример ответа

```json
{
  "status": "ok"
}
```

### Пример поиска (GET /get_logs)

```
GET /get_logs?level=info&limit=10
```

### Пример ответа

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

## 🔗 Документация

- Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Примеры curl:
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

## ⚙️ Конфигурация

| Флаг / Переменная | По умолчанию          | Описание                   |
|-------------------|-----------------------|----------------------------|
| LISTEN_ADDR       | :8080                 | Адрес для запуска API      |
| ELASTIC_URL       | http://localhost:9200 | URL Elasticsearch          |
| ELASTIC_USERNAME  | elastic               | Пользователь Elasticsearch |
| ELASTIC_PASSWORD  | change_me             | Пароль Elasticsearch       |
| LOG_LEVEL         | info                  | Уровень логирования        |
| JWT_SECRET        | changeme              | Секрет для JWT             |

---

## 🐳 Docker / Compose

1. Собрать образ:
   ```sh
   docker build -t log-stash-lite .
   ```
2. Запустить через docker-compose:
   ```sh
   docker-compose up --build
   ```
3. Пробросить параметры:
    - Через переменные окружения в docker-compose.yml
    - Можно монтировать volume для логов/конфига

---

## 🧪 Тестирование

- Запуск всех тестов:
  ```sh
  go test ./...
  ```
- Проверка покрытия:
  ```sh
  go test -cover ./...
  ```
- Интеграционные тесты:
  ```sh
  go test -tags=integration ./internal/api/integration
  ```

---

## 📝 FAQ / Примеры

- **Ошибка 401:** Проверьте JWT-токен, используйте /auth/token для генерации.
- **Elasticsearch не отвечает:** Проверьте docker-compose, переменные окружения.
- **Как добавить новый парсер?** См. ниже.
- **Как сменить хранилище?** Реализуйте интерфейс storage.Storage и передайте в Handler.

---

## ➕ Как добавить парсинг своего логгера

1. Создайте файл в `internal/parser/`, например, `mylogger.go`:
   ```go
   package parser

   func ParseMyLogger(log map[string]interface{}) (map[string]interface{}, error) {
       // Ваша логика парсинга
       return log, nil
   }
   ```
2. Зарегистрируйте парсер в `parser.go`:
   ```go
   // ...
   case "mylogger":
       parseFunc = ParseMyLogger
   // ...
   ```
3. Используйте `"parse_type": "mylogger"` в своих API-запросах.

---

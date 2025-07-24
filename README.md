# Log Stash Lite

Log Stash Lite is a lightweight service for storing and searching structured logs in Elasticsearch.

## Prerequisites

* [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed

## Running with Docker Compose

```
docker compose up
```

The service will listen on `http://localhost:8080` and start a local Elasticsearch instance on port `9200`.

Environment variables can be provided via a `.env` file or directly in `docker-compose.yml`. The defaults are:

| Variable          | Default                     | Description                          |
|-------------------|-----------------------------|--------------------------------------|
| `LISTEN_ADDR`     | `:8080`                     | HTTP listen address                  |
| `ELASTIC_URL`     | `http://localhost:9200`     | Elasticsearch HTTP endpoint          |
| `ELASTIC_USERNAME`| `elastic`                   | Elasticsearch username               |
| `ELASTIC_PASSWORD`| `change_me`                 | Elasticsearch password               |

## API Endpoints

| Method | Endpoint    | Description                                    |
|--------|-------------|------------------------------------------------|
| POST   | `/add_log`  | Index a single log entry.                      |
| POST   | `/add_logs` | Index multiple log entries.                    |
| GET    | `/get_logs` | Query logs using URL parameters as filters.    |
| GET    | `/logs_stats` | Retrieve statistics about stored logs.        |



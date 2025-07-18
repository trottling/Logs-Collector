APP_NAME := app
DOCKER_IMAGE := $(APP_NAME):latest
DOCKER_COMPOSE := docker compose
GO_BUILD := go build -o $(APP_NAME)

.PHONY: all build run logs stop down restart clean rebuild

all: build

build:
	@echo "🚧 Building Go binary..."
	$(GO_BUILD) .

docker-build:
	@echo "📦 Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

run:
	@echo "🚀 Starting containers..."
	$(DOCKER_COMPOSE) up -d

stop:
	@echo "⛔ Stopping containers..."
	$(DOCKER_COMPOSE) stop

down:
	@echo "🔥 Shutting down and removing containers..."
	$(DOCKER_COMPOSE) down -v

restart: down run

logs:
	@echo "📖 Logs:"
	$(DOCKER_COMPOSE) logs -f app

clean:
	@echo "🧽 Cleaning up..."
	rm -f $(APP_NAME)
	docker rmi $(DOCKER_IMAGE) || true

rebuild: clean build docker-build restart


SHELL := /bin/bash

.PHONY: up down build tidy gen test

up:
	docker compose -f infra/docker-compose.yml up -d

down:
	docker compose -f infra/docker-compose.yml down -v

tidy:
	go mod tidy

build:
	go build ./...

test:
	go test ./... -count=1

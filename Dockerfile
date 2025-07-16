# syntax=docker/dockerfile:1.4
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/server

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/app /app/app

ENV PORT=8080

EXPOSE 8080

CMD ["./app"]

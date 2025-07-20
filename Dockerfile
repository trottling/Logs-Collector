FROM golang:1.24-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

RUN chmod +x ./app

CMD ["./app"]

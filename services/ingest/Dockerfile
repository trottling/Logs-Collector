FROM golang:1.24 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/app ./services/ingest/cmd

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=build /out/app /app
EXPOSE 8002
ENTRYPOINT ["/app"]

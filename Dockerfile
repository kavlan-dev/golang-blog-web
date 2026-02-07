# Сборка
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app

# Запуск
FROM alpine
WORKDIR /app/
COPY --from=builder /app/app .
CMD ["./app"]

package main

import (
	"log"
	"net/http"

	"golang-blog-web/internal/config"
	"golang-blog-web/internal/handlers"
	"golang-blog-web/internal/middleware"
	"golang-blog-web/internal/services"
	"golang-blog-web/internal/storage/memory"
)

func main() {
	cfg := config.LoadConfig()

	storage := memory.New()
	service := services.New(storage)
	handler := handlers.New(service)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health/", handler.HealthCheck)
	mux.HandleFunc("POST /api/posts", handler.CreatePostHandler)
	mux.HandleFunc("GET /api/posts/", handler.GetPostHandler)
	mux.HandleFunc("PUT /api/posts", handler.UpdatePostHandler)
	mux.HandleFunc("DELETE /api/posts", handler.DeletePostHandler)

	log.Fatalf("Ошибка запуска сервера: %v", http.ListenAndServe(config.GetServerAddress(cfg), middleware.CORSMiddleware(cfg, mux)))
}

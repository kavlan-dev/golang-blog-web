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

	mainMux := http.NewServeMux()
	mainMux.HandleFunc("GET /health", handler.HealthCheck)
	mainMux.HandleFunc("GET /api/posts/", handler.GetPostHandler)

	secureMux := http.NewServeMux()
	secureMux.HandleFunc("POST /api/secure/posts", handler.CreatePostHandler)
	secureMux.HandleFunc("PUT /api/secure/posts/", handler.UpdatePostHandler)
	secureMux.HandleFunc("DELETE /api/secure/posts/", handler.DeletePostHandler)

	secureHandler := middleware.AuthMiddleware(cfg, secureMux)

	mainMux.Handle("/api/secure/", secureHandler)

	log.Fatalf("Ошибка запуска сервера: %v", http.ListenAndServe(config.GetServerAddress(cfg), middleware.CORSMiddleware(cfg, mainMux)))
}

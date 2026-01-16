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

	err := service.CreateFirstAdmin(cfg)
	if err != nil {
		log.Fatal("Не удалось создать администратора")
	}

	mainMux := http.NewServeMux()
	mainMux.HandleFunc("GET /health", handler.HealthCheck)
	mainMux.HandleFunc("GET /api/posts/", handler.GetPostHandler)
	mainMux.HandleFunc("POST /api/auth/register", handler.CreateUserHandler)

	authMux := http.NewServeMux()
	authMux.HandleFunc("POST /api/posts", handler.CreatePostHandler)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("PUT /api/secure/posts/", handler.UpdatePostHandler)
	adminMux.HandleFunc("DELETE /api/secure/posts/", handler.DeletePostHandler)
	adminMux.HandleFunc("PUT /api/secure/users/", handler.UpdateUserHandler)

	authHandler := middleware.AuthMiddleware(service, authMux)
	adminHandler := middleware.AuthAdminMiddleware(service, adminMux)

	combinedMux := http.NewServeMux()
	combinedMux.Handle("/", mainMux)
	combinedMux.Handle("POST /api/posts", authHandler)
	combinedMux.Handle("/api/secure/", adminHandler)

	log.Fatalf("Ошибка запуска сервера: %v", http.ListenAndServe(config.GetServerAddress(cfg), middleware.CORSMiddleware(cfg, combinedMux)))
}

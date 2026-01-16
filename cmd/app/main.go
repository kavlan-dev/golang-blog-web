package main

import (
	"net/http"

	"golang-blog-web/internal/config"
	"golang-blog-web/internal/handlers"
	"golang-blog-web/internal/middleware"
	"golang-blog-web/internal/services"
	"golang-blog-web/internal/storage/memory"
	"golang-blog-web/internal/utils"
)

func main() {
	log := utils.NewLogger()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error("Ошибка загрузки настроек", utils.Err(err))
		return
	}

	storage := memory.New()
	service := services.New(storage)
	handler := handlers.New(service, log)

	if err := service.CreateFirstAdmin(cfg); err != nil {
		log.Error("Не удалось создать администратора", utils.Err(err))
		return
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

	err = http.ListenAndServe(config.GetServerAddress(cfg), middleware.CORSMiddleware(cfg, combinedMux))
	if err != nil {
		log.Error("Ошибка запуска сервера", utils.Err(err))
		return
	}
}

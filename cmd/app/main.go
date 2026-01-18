package main

import (
	"go-blog-web/internal/config"
	"go-blog-web/internal/handlers"
	"go-blog-web/internal/middleware"
	"go-blog-web/internal/services"
	"go-blog-web/internal/storage/memory"
	"go-blog-web/internal/utils"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Ошибка загрузки настроек", err)
		return
	}
	log := utils.New(cfg.Env)
	log.Debug("Конфигурация сервиса", slog.Any("config", *cfg))
	log.Debug("Загрузка cors", slog.Any("Cors", cfg.Cors()))

	storage := memory.New()
	service := services.New(storage)
	handler := handlers.New(service, log)

	if err := service.CreateFirstAdmin(cfg); err != nil {
		log.Error("Не удалось создать администратора", utils.Err(err))
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health/", handler.HealthCheck)
	mux.HandleFunc("GET /api/posts/", handler.Posts)
	mux.HandleFunc("GET /api/posts/{id}/", handler.PostById)
	mux.HandleFunc("POST /api/auth/register/", handler.CreateUserHandler)

	mux.HandleFunc("POST /api/posts/", middleware.AuthMiddleware(service, handler.CreatePost))

	mux.HandleFunc("PUT /api/secure/posts/{id}/", middleware.AuthAdminMiddleware(service, handler.UpdatePost))
	mux.HandleFunc("DELETE /api/secure/posts/{id}/", middleware.AuthAdminMiddleware(service, handler.DeletePost))
	mux.HandleFunc("PUT /api/secure/users/{id}/", middleware.AuthAdminMiddleware(service, handler.UpdateUserHandler))

	err = http.ListenAndServe(cfg.ServerAddress(), middleware.CORSMiddleware(cfg, mux))
	if err != nil {
		log.Error("Ошибка запуска сервера", utils.Err(err))
		return
	}
}

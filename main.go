package main

import (
	"log"
	"net/http"

	"golang-blog-web/internal/config"
	"golang-blog-web/internal/handlers"
	"golang-blog-web/internal/middleware"
	"golang-blog-web/internal/services"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Инициализация сервиса для работы с записями
	postService := services.NewPostService()

	// Инициализация хендлера
	postHandler := handlers.NewPostHandler(postService)

	// Создаем основной маршрутизатор
	mux := http.NewServeMux()

	// Настройка маршрутов
	mux.HandleFunc("/api/posts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postHandler.CreatePostHandler(w, r)
		case http.MethodGet:
			postHandler.GetPostHandler(w, r)
		case http.MethodPut:
			postHandler.UpdatePostHandler(w, r)
		case http.MethodDelete:
			postHandler.DeletePostHandler(w, r)
		}
	})

	// Запуск сервера с CORS middleware
	log.Printf("Запуск сервера на порту %s...", cfg.ServerPort)
	log.Printf("API блога доступен по адресу: http://%s", cfg.GetServerAddress())

	// Оборачиваем весь маршрутизатор в CORS middleware и запускаем сервер
	err := http.ListenAndServe(cfg.GetServerAddress(), middleware.CORSMiddleware(cfg, mux))
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

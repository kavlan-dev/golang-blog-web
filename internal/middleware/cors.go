package middleware

import (
	"golang-blog-web/internal/config"
	"net/http"
)

// CORSMiddleware добавляет заголовки CORS к ответам
func CORSMiddleware(cfg *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовки CORS
		w.Header().Set("Access-Control-Allow-Origin", cfg.CORSAllowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, , Authorization")

		// Обработка OPTIONS запросов (preflight)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}

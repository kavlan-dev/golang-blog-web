package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type response map[string]any

type servicesInterface interface {
	postsService
	userService
}

type handler struct {
	service servicesInterface
	log     *slog.Logger
}

func New(service servicesInterface, log *slog.Logger) *handler {
	return &handler{
		service: service,
		log:     log,
	}
}

func (h *handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		"message": "Сервис работает исправно",
	})
}

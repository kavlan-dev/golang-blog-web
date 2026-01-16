package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response map[string]any

type ServicesInterface interface {
	PostsService
	UserService
}

type Handler struct {
	service ServicesInterface
	log     *slog.Logger
}

func New(service ServicesInterface, log *slog.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		"message": "Сервис работает исправно",
	})
}

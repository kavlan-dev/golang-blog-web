package handlers

import (
	"encoding/json"
	"net/http"
)

type Response map[string]any

type ServicesInterface interface {
	PostsService
}

type Handler struct {
	service ServicesInterface
}

func New(service ServicesInterface) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		"message": "Сервис работает исправно",
	})
}

package handlers

import (
	"encoding/json"
	"net/http"
)

type Response map[string]any

type ServicesInterface interface {
	PostsService
}

type PostHandler struct {
	service ServicesInterface
}

func New(service ServicesInterface) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		"message": "Сервис работает исправно",
	})
}

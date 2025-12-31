package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang-blog-web/internal/services"
)

// PostHandler обрабатывает HTTP запросы для работы с записями блога
type PostHandler struct {
	service *services.PostService
}

// NewPostHandler создает новый экземпляр хендлера
func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

// CreatePostHandler обрабатывает создание новой записи
func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Ошибка декодирования JSON: %v", err)
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	post, err := h.service.CreatePost(request.Title, request.Content)
	if err != nil {
		log.Printf("Ошибка создания записи: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

// GetPostHandler возвращает все записи или запись по id
func (h *PostHandler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/posts/"):]
	if idStr == "" {
		posts := h.service.GetAllPosts()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	post, err := h.service.GetPostByID(id)
	if err != nil {
		log.Printf("Ошибка получения записи: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// UpdatePostHandler обновляет запись
func (h *PostHandler) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/posts/"):]
	if idStr == "" {
		http.Error(w, "отсутствует id записи", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	var request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Ошибка декодирования JSON: %v", err)
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	post, err := h.service.UpdatePost(id, request.Title, request.Content)
	if err != nil {
		log.Printf("Ошибка обновления записи: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// DeletePostHandler удаляет запись
func (h *PostHandler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/posts/"):]
	if idStr == "" {
		http.Error(w, "отсутствует параметр id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	err = h.service.DeletePost(id)
	if err != nil {
		log.Printf("Ошибка удаления записи: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

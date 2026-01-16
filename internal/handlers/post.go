package handlers

import (
	"encoding/json"
	"golang-blog-web/internal/models"
	"net/http"
	"strconv"
)

type PostsService interface {
	CreatePost(newPost *models.Post) error
	GetAllPosts() *[]models.Post
	GetPostByID(id uint) (*models.Post, error)
	UpdatePost(id uint, updatePost *models.Post) error
	DeletePost(id uint) error
}

func (h *Handler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req models.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	newPost := &models.Post{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.service.CreatePost(newPost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func (h *Handler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
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

	post, err := h.service.GetPostByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *Handler) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/secure/posts/"):]
	if idStr == "" {
		http.Error(w, "отсутствует id записи", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	var req models.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	updatePost := &models.Post{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.service.UpdatePost(uint(id), updatePost); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatePost)
}

func (h *Handler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/secure/posts/"):]
	if idStr == "" {
		http.Error(w, "отсутствует параметр id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	err = h.service.DeletePost(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

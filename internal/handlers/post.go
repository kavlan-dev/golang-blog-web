package handlers

import (
	"encoding/json"
	"golang-blog-web/internal/models"
	"golang-blog-web/internal/utils"
	"log/slog"
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
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req models.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Ошибка в теле запроса", utils.Err(err))
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	newPost := &models.Post{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.service.CreatePost(newPost); err != nil {
		h.log.Error("Ошибка при создании записи", utils.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.log.Info("Успешно создана запись", slog.Int("post id", int(newPost.ID)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func (h *Handler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/posts/"):]
	if idStr == "" {
		posts := h.service.GetAllPosts()

		h.log.Info("Получены все записи")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("Не верный ввод id", utils.Err(err))
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	post, err := h.service.GetPostByID(uint(id))
	if err != nil {
		h.log.Error("Ошибка при попытке получить запись", utils.Err(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.log.Info("Успешно получена запись", slog.String("post id", idStr))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *Handler) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/secure/posts/"):]
	if idStr == "" {
		h.log.Error("Отсутствует id")
		http.Error(w, "отсутствует id записи", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("Не верный ввод id", utils.Err(err))
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	var req models.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Ошибка в теле запроса", utils.Err(err))
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	updatePost := &models.Post{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.service.UpdatePost(uint(id), updatePost); err != nil {
		h.log.Error("Ошибка при обновлении записи", utils.Err(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.log.Info("Успешно обновлена запись", slog.String("post id", idStr))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatePost)
}

func (h *Handler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/secure/posts/"):]
	if idStr == "" {
		h.log.Error("Отсутствует id")
		http.Error(w, "отсутствует параметр id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("Не верный ввод id", utils.Err(err))
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	err = h.service.DeletePost(uint(id))
	if err != nil {
		h.log.Error("Ошибка при удалении записи", utils.Err(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.log.Info("Успешно удалена запись", slog.String("post id", idStr))
	w.WriteHeader(http.StatusNoContent)
}

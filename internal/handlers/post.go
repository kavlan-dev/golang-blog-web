package handlers

import (
	"encoding/json"
	"go-blog-web/internal/models"
	"go-blog-web/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
)

type postsService interface {
	CreatePost(newPost *models.Post) error
	AllPosts() *[]models.Post
	PostByID(id uint) (*models.Post, error)
	PostByTitle(title string) (*models.Post, error)
	UpdatePost(id uint, updatePost *models.Post) error
	DeletePost(id uint) error
}

func (h *handler) CreatePost(w http.ResponseWriter, r *http.Request) {
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
		Tags:    req.Tags,
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

// TODO Реализовать сортировку и фильтрацию
func (h *handler) Posts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	posts := h.service.AllPosts()

	if len(*posts) == 0 {
		h.log.Warn("Список записей пуст")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response{"message": "Записи отсутствуют"})
		return
	}

	h.log.Info("Получены все записи", slog.Int("count", len(*posts)))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *handler) PostById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("Не верный ввод id", utils.Err(err))
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	post, err := h.service.PostByID(uint(id))
	if err != nil {
		h.log.Error("Ошибка при попытке получить запись", utils.Err(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.log.Info("Успешно получена запись", slog.Int("post id", id))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *handler) PostByTitle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	title := r.PathValue("title")
	if title == "" {
		h.log.Error("Не указан заголовок")
		http.Error(w, "Заголовок не может быть пустым", http.StatusBadRequest)
		return
	}

	post, err := h.service.PostByTitle(title)
	if err != nil {
		h.log.Error("Ошибка при попытке получить запись", utils.Err(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.log.Info("Успешно получена запись", slog.String("post title", title))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
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
		Tags:    req.Tags,
	}

	if err := h.service.UpdatePost(uint(id), updatePost); err != nil {
		h.log.Error("Ошибка при обновлении записи", utils.Err(err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.log.Info("Успешно обновлена запись", slog.Int("post id", id))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatePost)
}

func (h *handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
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

	h.log.Info("Успешно удалена запись", slog.Int("post id", id))
	w.WriteHeader(http.StatusNoContent)
}

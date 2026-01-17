package handlers

import (
	"encoding/json"
	"go-blog-web/internal/models"
	"go-blog-web/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
)

type UserService interface {
	CreateUser(newUser *models.User) error
	UpdateUser(id uint, updateUser *models.User) error
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Ошибка в теле запроса", utils.Err(err))
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	newUser := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     "user",
		// Role: req.Role,
	}
	if err := h.service.CreateUser(newUser); err != nil {
		h.log.Error("Ошибка при создании пользователя", utils.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.log.Info("Успешно создан пользователь", slog.Int("user id", int(newUser.ID)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Ошибка в теле запроса", utils.Err(err))
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	updateUser := &models.User{
		Role: req.Role,
	}

	if err := h.service.UpdateUser(uint(id), updateUser); err != nil {
		h.log.Error("Ошибка при обновлении пользователя", utils.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.log.Info("Успешно обновлен пользователь", slog.Int("user id", id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updateUser)
}

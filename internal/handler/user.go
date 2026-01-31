package handler

import (
	"encoding/json"
	"go-blog-web/internal/model"
	"go-blog-web/internal/util"
	"log/slog"
	"net/http"
	"strconv"
)

type userService interface {
	CreateUser(newUser *model.User) error
	UpdateUser(id uint, updateUser *model.User) error
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Ошибка в теле запроса", util.Err(err))
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	newUser := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     "user",
		// Role: req.Role,
	}
	if err := h.service.CreateUser(newUser); err != nil {
		h.log.Error("Ошибка при создании пользователя", util.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.log.Info("Успешно создан пользователь", slog.Int("user id", int(newUser.ID)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.log.Warn("Использован не подходящий метод", slog.String("method", r.Method))
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("Не верный ввод id", util.Err(err))
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	var req model.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Ошибка в теле запроса", util.Err(err))
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	updateUser := &model.User{
		Role: req.Role,
	}

	if err := h.service.UpdateUser(uint(id), updateUser); err != nil {
		h.log.Error("Ошибка при обновлении пользователя", util.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.log.Info("Успешно обновлен пользователь", slog.Int("user id", id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updateUser)
}

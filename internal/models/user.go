package models

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"` // user, admin
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	// Role     string `json:"role"`
}

type UpdateUserRequest struct {
	Role string `json:"role"`
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Username) == "" {
		return fmt.Errorf("Имя пользователя не может быть пустым")
	}
	if strings.TrimSpace(u.Password) == "" {
		return fmt.Errorf("Пароль не может быть пустым")
	}
	if strings.TrimSpace(u.Email) == "" {
		return fmt.Errorf("Укажите почту")
	}
	if !strings.Contains(strings.TrimSpace(u.Email), "@") {
		return fmt.Errorf("Некорректный формат email")
	}
	return nil
}

func (u *User) IsUserUnique(users []User) bool {
	for id, user := range users {
		if id == int(u.ID) {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(user.Username), strings.TrimSpace(u.Username)) {
			return false
		}
		if strings.EqualFold(strings.TrimSpace(user.Email), strings.TrimSpace(u.Email)) {
			return false
		}
	}
	return true
}

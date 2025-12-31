package models

import (
	"errors"
	"strings"
	"time"
)

// Post представляет модель записи блога
type Post struct {
	ID        int       `json:"id"`         // Уникальный идентификатор записи
	Title     string    `json:"title"`      // Заголовок записи
	Content   string    `json:"content"`    // Содержимое записи
	CreatedAt time.Time `json:"created_at"` // Дата и время создания
	UpdatedAt time.Time `json:"updated_at"` // Дата и время последнего обновления
}

const (
	MaxTitleLength   = 500
	MaxContentLength = 10000
)

// Validate проверяет корректность данных записи
func (p *Post) Validate() error {
	if strings.TrimSpace(p.Title) == "" {
		return errors.New("заголовок не может быть пустым")
	}
	if len(p.Title) > MaxTitleLength {
		return errors.New("заголовок слишком длинный")
	}
	if strings.TrimSpace(p.Content) == "" {
		return errors.New("контент не может быть пустым")
	}
	if len(p.Content) > MaxContentLength {
		return errors.New("контент слишком длинный")
	}
	return nil
}

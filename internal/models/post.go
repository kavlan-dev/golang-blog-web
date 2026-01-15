package models

import (
	"fmt"
	"strings"
	"time"
)

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *Post) Validate() error {
	if strings.TrimSpace(p.Title) == "" {
		return fmt.Errorf("заголовок не может быть пустым")
	}
	if strings.TrimSpace(p.Content) == "" {
		return fmt.Errorf("контент не может быть пустым")
	}
	return nil
}

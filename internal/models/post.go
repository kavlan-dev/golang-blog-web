package models

import (
	"fmt"
	"strings"
	"time"
)

// type Comment struct {
// 	ID        uint      `json:"id"`
// 	Content   string    `json:"content"`
// 	Author    string    `json:"author"`
// 	CreatedAt time.Time `json:"created_at"`
// }

type Post struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags,omitempty"`
	// Comments  []Comment `json:"comments,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags,omitempty"`
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

func (p *Post) IsTitleUnique(posts []Post) bool {
	for id, post := range posts {
		if id == int(p.ID) {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(post.Title), strings.TrimSpace(p.Title)) {
			return false
		}
	}
	return true
}

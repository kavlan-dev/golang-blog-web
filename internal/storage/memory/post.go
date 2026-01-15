package memory

import (
	"errors"
	"fmt"
	"golang-blog-web/internal/models"
	"strings"
	"time"
)

func (s *Storage) isTitleUnique(title string, excludeID uint) bool {
	for id, post := range s.posts {
		if id != excludeID && strings.EqualFold(strings.TrimSpace(post.Title), strings.TrimSpace(title)) {
			return false
		}
	}
	return true
}

func (s *Storage) CreatePost(newPost *models.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ok := s.isTitleUnique(newPost.Title, newPost.ID)
	if ok != true {
		return fmt.Errorf("Запись с таким заголовком уже существует")
	}

	newPost.ID = s.nextID
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()

	s.posts[s.nextID] = newPost
	s.nextID++

	return nil
}

func (s *Storage) FindPosts() *[]models.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	allPosts := make([]models.Post, 0, len(s.posts))
	for _, post := range s.posts {
		allPosts = append(allPosts, *post)
	}

	return &allPosts
}

func (s *Storage) FindPostBiId(id uint) (*models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[id]
	if !exists {
		return nil, fmt.Errorf("запись с id %d не найдена", id)
	}

	return post, nil
}

func (s *Storage) UpdatePost(id uint, updatePost *models.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[id]
	if !exists {
		return fmt.Errorf("запись с id %d не найдена", id)
	}

	if !s.isTitleUnique(post.Title, id) {
		return errors.New("запись с таким заголовком уже существует")
	}

	post.Title = updatePost.Title
	post.Content = updatePost.Content
	post.UpdatedAt = time.Now()

	s.posts[id] = post

	return nil
}

func (s *Storage) DeletePost(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.posts[id]; !exists {
		return fmt.Errorf("запись с id %d не найдена", id)
	}

	delete(s.posts, id)
	return nil
}

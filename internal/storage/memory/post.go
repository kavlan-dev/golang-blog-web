package memory

import (
	"errors"
	"fmt"
	"go-blog-web/internal/models"
	"time"
)

func (s *storage) CreatePost(newPost *models.Post) error {
	posts := s.FindPosts()
	s.mu.Lock()
	defer s.mu.Unlock()

	newPost.ID = s.nextPostId
	if !newPost.IsTitleUnique(*posts) {
		return fmt.Errorf("Запись с таким заголовком уже существует")
	}

	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()

	s.posts[s.nextPostId] = newPost
	s.nextPostId++

	return nil
}

func (s *storage) FindPosts() *[]models.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	allPosts := make([]models.Post, 0, len(s.posts))
	for _, post := range s.posts {
		allPosts = append(allPosts, *post)
	}

	return &allPosts
}

func (s *storage) FindPostById(id uint) (*models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[id]
	if !exists {
		return nil, fmt.Errorf("запись с id %d не найдена", id)
	}

	return post, nil
}

func (s *storage) FindPostByTitle(title string) (*models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, post := range s.posts {
		if post.Title == title {
			return post, nil
		}
	}

	return nil, fmt.Errorf("запись с заголовком %s не найдена", title)
}

func (s *storage) UpdatePost(id uint, updatePost *models.Post) error {
	posts := s.FindPosts()
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[id]
	if !exists {
		return fmt.Errorf("запись с id %d не найдена", id)
	}

	if !post.IsTitleUnique(*posts) {
		return errors.New("запись с таким заголовком уже существует")
	}

	post.Title = updatePost.Title
	post.Content = updatePost.Content
	post.Tags = updatePost.Tags
	post.UpdatedAt = time.Now()

	s.posts[id] = post

	return nil
}

func (s *storage) DeletePost(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.posts[id]; !exists {
		return fmt.Errorf("запись с id %d не найдена", id)
	}

	delete(s.posts, id)
	return nil
}

package services

import "go-blog-web/internal/models"

type PostsStorage interface {
	CreatePost(newPost *models.Post) error
	FindPosts() *[]models.Post
	FindPostById(id uint) (*models.Post, error)
	UpdatePost(id uint, updatePost *models.Post) error
	DeletePost(id uint) error
}

func (s *Service) CreatePost(newPost *models.Post) error {
	if err := newPost.Validate(); err != nil {
		return err
	}

	return s.storage.CreatePost(newPost)
}

func (s *Service) AllPosts() *[]models.Post {
	return s.storage.FindPosts()
}

func (s *Service) PostByID(id uint) (*models.Post, error) {
	return s.storage.FindPostById(id)
}

func (s *Service) UpdatePost(id uint, updatePost *models.Post) error {
	if err := updatePost.Validate(); err != nil {
		return err
	}

	return s.storage.UpdatePost(id, updatePost)
}

func (s *Service) DeletePost(id uint) error {
	return s.storage.DeletePost(id)
}

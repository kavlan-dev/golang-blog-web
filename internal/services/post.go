package services

import "go-blog-web/internal/models"

type postsStorage interface {
	CreatePost(newPost *models.Post) error
	FindPosts() *[]models.Post
	FindPostById(id uint) (*models.Post, error)
	FindPostByTitle(title string) (*models.Post, error)
	UpdatePost(id uint, updatePost *models.Post) error
	DeletePost(id uint) error
}

func (s *service) CreatePost(newPost *models.Post) error {
	if err := newPost.Validate(); err != nil {
		return err
	}

	return s.storage.CreatePost(newPost)
}

func (s *service) AllPosts() *[]models.Post {
	return s.storage.FindPosts()
}

func (s *service) PostByID(id uint) (*models.Post, error) {
	return s.storage.FindPostById(id)
}

func (s *service) PostByTitle(title string) (*models.Post, error) {
	return s.storage.FindPostByTitle(title)
}

func (s *service) UpdatePost(id uint, updatePost *models.Post) error {
	if err := updatePost.Validate(); err != nil {
		return err
	}

	return s.storage.UpdatePost(id, updatePost)
}

func (s *service) DeletePost(id uint) error {
	return s.storage.DeletePost(id)
}

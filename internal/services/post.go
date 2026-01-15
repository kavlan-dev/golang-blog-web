package services

import (
	"golang-blog-web/internal/models"
)

type PostsStorage interface {
	CreatePost(newPost *models.Post) error
	FindPosts() *[]models.Post
	FindPostBiId(id uint) (*models.Post, error)
	UpdatePost(id uint, updatePost *models.Post) error
	DeletePost(id uint) error
}

func (s *PostService) CreatePost(newPost *models.Post) error {
	if err := newPost.Validate(); err != nil {
		return err
	}

	return s.storage.CreatePost(newPost)
}

func (s *PostService) GetAllPosts() *[]models.Post {
	return s.storage.FindPosts()
}

func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	return s.storage.FindPostBiId(id)
}

func (s *PostService) UpdatePost(id uint, updatePost *models.Post) error {
	if err := updatePost.Validate(); err != nil {
		return err
	}

	return s.storage.UpdatePost(id, updatePost)
}

func (s *PostService) DeletePost(id uint) error {
	return s.storage.DeletePost(id)
}

package services

import (
	"fmt"
	"go-blog-web/internal/config"
	"go-blog-web/internal/models"
)

type UsersStorage interface {
	CreateUser(user *models.User) error
	UserByUsername(username string) (*models.User, error)
	UpdateUser(id uint, updateUser *models.User) error
}

func (s *Service) CreateUser(newUser *models.User) error {
	if err := newUser.Validate(); err != nil {
		return err
	}

	return s.storage.CreateUser(newUser)
}

func (s *Service) userByUsername(username string) (*models.User, error) {
	return s.storage.UserByUsername(username)
}

func (s *Service) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := s.userByUsername(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, fmt.Errorf("Не верный пароль")
	}

	return user, nil
}

func (s *Service) CreateFirstAdmin(cfg *config.Config) error {
	return s.storage.CreateUser(&models.User{
		Username: cfg.Admin.Username,
		Password: cfg.Admin.Password,
		Email:    cfg.Admin.Email,
		Role:     "admin",
	})
}

func (s *Service) UpdateUser(id uint, updateUser *models.User) error {
	if err := updateUser.Validate(); err != nil {
		return err
	}

	return s.UpdateUser(id, updateUser)
}

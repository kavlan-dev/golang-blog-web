package services

import (
	"fmt"
	"go-blog-web/internal/config"
	"go-blog-web/internal/models"
)

type UsersStorage interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(id uint, updateUser *models.User) error
}

func (s *Service) CreateUser(newUser *models.User) error {
	if err := newUser.Validate(); err != nil {
		return err
	}

	return s.storage.CreateUser(newUser)
}

func (s *Service) getUserByUsername(username string) (*models.User, error) {
	return s.storage.GetUserByUsername(username)
}

func (s *Service) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := s.getUserByUsername(username)
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
		Username: cfg.AdminUsername,
		Password: cfg.AdminPassword,
		Email:    cfg.AdminEmail,
		Role:     "admin",
	})
}

func (s *Service) UpdateUser(id uint, updateUser *models.User) error {
	if err := updateUser.Validate(); err != nil {
		return err
	}

	return s.UpdateUser(id, updateUser)
}

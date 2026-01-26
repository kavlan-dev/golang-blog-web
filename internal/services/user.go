package services

import (
	"encoding/hex"
	"fmt"
	"go-blog-web/internal/config"
	"go-blog-web/internal/models"
	"go-blog-web/internal/utils"
)

type usersStorage interface {
	CreateUser(user *models.User) error
	UserByUsername(username string) (*models.User, error)
	UpdateUser(id uint, updateUser *models.User) error
}

func (s *service) CreateUser(newUser *models.User) error {
	if err := newUser.Validate(); err != nil {
		return err
	}

	newUser.Password = utils.HashPassword(newUser.Password)

	return s.storage.CreateUser(newUser)
}

func (s *service) userByUsername(username string) (*models.User, error) {
	return s.storage.UserByUsername(username)
}

func (s *service) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := s.userByUsername(username)
	if err != nil {
		return nil, err
	}

	storedHash := user.Password
	if len(storedHash) < 64 {
		return nil, fmt.Errorf("некорректный формат хэша пароля")
	}

	saltHex := storedHash[:32]
	salt, err := hex.DecodeString(saltHex)
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования соли")
	}

	hashedPassword := utils.HashPasswordWithSalt(password, salt)
	if user.Password != hashedPassword {
		return nil, fmt.Errorf("Не верный пароль")
	}

	return user, nil
}

// Создает первого администратора для последующего использования эндпоинтов доступных только администратору
// Данные для входа устанавливаются в конфигурационном файле JSON
func (s *service) CreateFirstAdmin(cfg *config.Config) error {
	admin := &models.User{
		Username: cfg.Admin.Username,
		Password: utils.HashPassword(cfg.Admin.Password),
		Email:    cfg.Admin.Email,
		Role:     "admin",
	}
	return s.storage.CreateUser(admin)
}

func (s *service) UpdateUser(id uint, updateUser *models.User) error {
	if err := updateUser.Validate(); err != nil {
		return err
	}

	return s.storage.UpdateUser(id, updateUser)
}

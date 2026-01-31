package service

import (
	"encoding/hex"
	"fmt"
	"go-blog-web/internal/config"
	"go-blog-web/internal/model"
	"go-blog-web/internal/util"
)

type userStorage interface {
	CreateUser(user *model.User) error
	UserByUsername(username string) (*model.User, error)
	UpdateUser(id uint, updateUser *model.User) error
}

func (s *service) CreateUser(newUser *model.User) error {
	if err := newUser.Validate(); err != nil {
		return err
	}

	newUser.Password = util.HashPassword(newUser.Password)

	return s.storage.CreateUser(newUser)
}

func (s *service) userByUsername(username string) (*model.User, error) {
	return s.storage.UserByUsername(username)
}

func (s *service) AuthenticateUser(username, password string) (*model.User, error) {
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

	hashedPassword := util.HashPasswordWithSalt(password, salt)
	if user.Password != hashedPassword {
		return nil, fmt.Errorf("Не верный пароль")
	}

	return user, nil
}

// Создает первого администратора для последующего использования эндпоинтов доступных только администратору
// Данные для входа устанавливаются в конфигурационном файле JSON
func (s *service) CreateFirstAdmin(cfg *config.Config) error {
	admin := &model.User{
		Username: cfg.Admin.Username,
		Password: util.HashPassword(cfg.Admin.Password),
		Email:    cfg.Admin.Email,
		Role:     "admin",
	}
	return s.storage.CreateUser(admin)
}

func (s *service) UpdateUser(id uint, updateUser *model.User) error {
	if err := updateUser.Validate(); err != nil {
		return err
	}

	return s.storage.UpdateUser(id, updateUser)
}

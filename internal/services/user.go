package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
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

	newUser.Password = hashPassword(newUser.Password)

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

	storedHash := user.Password
	if len(storedHash) < 64 {
		return nil, fmt.Errorf("некорректный формат хэша пароля")
	}

	saltHex := storedHash[:32]
	salt, err := hex.DecodeString(saltHex)
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования соли")
	}

	hashedPassword := hashPasswordWithSalt(password, salt)
	if user.Password != hashedPassword {
		return nil, fmt.Errorf("Не верный пароль")
	}

	return user, nil
}

// Создает первого администратора для последующего использования эндпоинтов доступных только администратору
// Данные для входа устанавливаются в конфигурационном файле JSON
func (s *Service) CreateFirstAdmin(cfg *config.Config) error {
	admin := &models.User{
		Username: cfg.Admin.Username,
		Password: hashPassword(cfg.Admin.Password),
		Email:    cfg.Admin.Email,
		Role:     "admin",
	}
	return s.storage.CreateUser(admin)
}

func (s *Service) UpdateUser(id uint, updateUser *models.User) error {
	if err := updateUser.Validate(); err != nil {
		return err
	}

	return s.storage.UpdateUser(id, updateUser)
}

func hashPassword(password string) string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return hashPasswordWithSalt(password, []byte{})
	}
	return hashPasswordWithSalt(password, salt)
}

func hashPasswordWithSalt(password string, salt []byte) string {
	h := sha256.New()
	h.Write(salt)
	h.Write([]byte(password))
	hashed := h.Sum(nil)

	result := make([]byte, len(salt)+len(hashed))
	copy(result, salt)
	copy(result[len(salt):], hashed)

	return hex.EncodeToString(result)
}

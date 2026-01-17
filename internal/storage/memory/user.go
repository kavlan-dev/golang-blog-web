package memory

import (
	"fmt"
	"go-blog-web/internal/models"
	"strings"
	"time"
)

func (s *Storage) isUserUnique(username, email string, excludeID uint) bool {
	for id, user := range s.users {
		if id != excludeID && strings.EqualFold(strings.TrimSpace(user.Username), strings.TrimSpace(username)) && strings.EqualFold(strings.TrimSpace(user.Email), strings.TrimSpace(email)) {
			return false
		}
	}
	return true
}

func (s *Storage) CreateUser(newUser *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	newUser.ID = s.nextUserId
	ok := s.isUserUnique(newUser.Username, newUser.Email, newUser.ID)
	if !ok {
		return fmt.Errorf("")
	}

	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	s.users[newUser.ID] = newUser
	s.nextUserId++

	return nil
}

func (s *Storage) GetUserByUsername(username string) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("Пользователь не найден")
}

func (s *Storage) UpdateUser(id uint, updateUser *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return fmt.Errorf("пользователь с id %d не найден", id)
	}

	if !s.isUserUnique(user.Username, user.Email, id) {
		return fmt.Errorf("пользователь с таким именем или почтой уже существует")
	}

	user.Role = updateUser.Role
	user.UpdatedAt = time.Now()

	s.users[id] = user

	return nil
}

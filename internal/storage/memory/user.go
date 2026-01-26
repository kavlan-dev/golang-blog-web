package memory

import (
	"fmt"
	"go-blog-web/internal/models"
	"time"
)

func (s *storage) FindUsers() *[]models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	allUsers := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		allUsers = append(allUsers, *user)
	}

	return &allUsers
}

func (s *storage) CreateUser(newUser *models.User) error {
	users := s.FindUsers()
	s.mu.Lock()
	defer s.mu.Unlock()

	newUser.ID = s.nextUserId
	if !newUser.IsUserUnique(*users) {
		return fmt.Errorf("пользователь с таким именем или почтой уже существует")
	}

	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	s.users[newUser.ID] = newUser
	s.nextUserId++

	return nil
}

func (s *storage) UserByUsername(username string) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("Пользователь не найден")
}

func (s *storage) UpdateUser(id uint, updateUser *models.User) error {
	users := s.FindUsers()
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return fmt.Errorf("пользователь с id %d не найден", id)
	}

	if !user.IsUserUnique(*users) {
		return fmt.Errorf("пользователь с таким именем или почтой уже существует")
	}

	user.Role = updateUser.Role
	user.UpdatedAt = time.Now()

	s.users[id] = user

	return nil
}

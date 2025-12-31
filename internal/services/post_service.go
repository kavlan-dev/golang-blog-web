package services

import (
	"errors"
	"fmt"
	"golang-blog-web/internal/models"
	"log"
	"strings"
	"sync"
	"time"
)

// PostService управляет операциями с записями блога
type PostService struct {
	posts  map[int]models.Post
	mu     sync.Mutex
	nextID int
}

// NewPostService создает новый экземпляр сервиса для работы с записями
func NewPostService() *PostService {
	return &PostService{
		posts:  make(map[int]models.Post),
		nextID: 1,
	}
}

// isTitleUnique проверяет заголовок на уникальность
func (s *PostService) isTitleUnique(title string, excludeID int) bool {
	for id, post := range s.posts {
		if id != excludeID && strings.EqualFold(strings.TrimSpace(post.Title), strings.TrimSpace(title)) {
			return false
		}
	}
	return true
}

// CreatePost создает новую запись в блоге
func (s *PostService) CreatePost(title, content string) (models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post := models.Post{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Валидация данных
	if err := post.Validate(); err != nil {
		return models.Post{}, err
	}

	// Проверка на уникальность заголовка
	if !s.isTitleUnique(title, 0) {
		return models.Post{}, errors.New("запись с таким заголовком уже существует")
	}

	post.ID = s.nextID
	s.posts[post.ID] = post
	s.nextID++

	log.Printf("Создана новая запись с ID: %d", post.ID)
	return post, nil
}

// GetAllPosts возвращает все записи блога
func (s *PostService) GetAllPosts() []models.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	allPosts := make([]models.Post, 0, len(s.posts))
	for _, post := range s.posts {
		allPosts = append(allPosts, post)
	}

	log.Printf("Получено %d записей", len(allPosts))
	return allPosts
}

// GetPostByID возвращает запись по идентификатору
func (s *PostService) GetPostByID(id int) (models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[id]
	if !exists {
		return models.Post{}, fmt.Errorf("запись с id %d не найдена", id)
	}

	log.Printf("Получена запись с ID: %d", id)
	return post, nil
}

// UpdatePost обновляет существующую запись
func (s *PostService) UpdatePost(id int, title, content string) (models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[id]
	if !exists {
		return models.Post{}, fmt.Errorf("запись с id %d не найдена", id)
	}

	// Проверка на уникальность заголовка
	if !s.isTitleUnique(title, id) {
		return models.Post{}, errors.New("запись с таким заголовком уже существует")
	}

	// Валидация данных
	if err := post.Validate(); err != nil {
		return models.Post{}, err
	}

	// Обновляем данные
	post.Title = title
	post.Content = content
	post.UpdatedAt = time.Now()

	s.posts[id] = post
	log.Printf("Обновлена запись с ID: %d", id)
	return post, nil
}

// DeletePost удаляет запись по идентификатору
func (s *PostService) DeletePost(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.posts[id]; !exists {
		return fmt.Errorf("запись с id %d не найдена", id)
	}

	delete(s.posts, id)
	log.Printf("Удалена запись с ID: %d", id)
	return nil
}

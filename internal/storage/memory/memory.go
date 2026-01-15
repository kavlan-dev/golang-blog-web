package memory

import (
	"golang-blog-web/internal/models"
	"sync"
)

type Storage struct {
	posts  map[uint]*models.Post
	mu     sync.Mutex
	nextID uint
}

func New() *Storage {
	return &Storage{
		posts: make(map[uint]*models.Post),
	}
}

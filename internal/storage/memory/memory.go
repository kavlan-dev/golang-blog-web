package memory

import (
	"golang-blog-web/internal/models"
	"sync"
)

type Storage struct {
	posts      map[uint]*models.Post
	users      map[uint]*models.User
	mu         sync.Mutex
	nextPostId uint
	nextUserId uint
}

func New() *Storage {
	return &Storage{
		posts: make(map[uint]*models.Post),
		users: make(map[uint]*models.User),
	}
}

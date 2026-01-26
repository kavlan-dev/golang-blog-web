package memory

import (
	"go-blog-web/internal/models"
	"sync"
)

type storage struct {
	posts      map[uint]*models.Post
	users      map[uint]*models.User
	mu         sync.Mutex
	nextPostId uint
	nextUserId uint
}

func New() *storage {
	return &storage{
		posts: make(map[uint]*models.Post),
		users: make(map[uint]*models.User),
		mu:    sync.Mutex{},
	}
}

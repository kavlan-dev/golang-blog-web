package services

type storageInterface interface {
	postsStorage
	usersStorage
}

type service struct {
	storage storageInterface
}

func New(storage storageInterface) *service {
	return &service{storage: storage}
}

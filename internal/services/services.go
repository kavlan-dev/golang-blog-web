package services

type StorageInterface interface {
	PostsStorage
}

type PostService struct {
	storage StorageInterface
}

func New(storage StorageInterface) *PostService {
	return &PostService{storage: storage}
}

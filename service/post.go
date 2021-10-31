package service

import (
	"errors"
	"go-rest-api/entity"
	"go-rest-api/repository"
	"math/rand"
)

var (
	postRepository repository.PostRepository
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

func NewPostService(repository repository.PostRepository) PostService {
	postRepository = repository
	return &Service{}
}

type Service struct{}

func (*Service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("The post is empty")
	}

	if post.Title == "" {
		return errors.New("The post title is empty")
	}

	return nil
}

func (*Service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()
	return postRepository.Save(post)
}

func (*Service) FindAll() ([]entity.Post, error) {
	return postRepository.FindAll()
}

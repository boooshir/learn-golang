package service

import (
	"go-filed-errors/internals/model"
	"go-filed-errors/internals/repository"

	"github.com/go-playground/validator/v10"
)

type PostService interface {
	CreatePost(post *model.Post) (*model.Post, error)
}

type postServiceImpl struct {
	repo repository.PostRepository
}

func NewPostservice(repository repository.PostRepository) PostService {
	return &postServiceImpl{repo: repository}
}

func (service *postServiceImpl) CreatePost(post *model.Post) (*model.Post, error) {
	validate := validator.New()
	if err := validate.Struct(post); err != nil {
		return nil, err
	}
	// save the repository
	if err := service.repo.CreatePost(post); err != nil {
		return nil, err
	}
	return post, nil
}

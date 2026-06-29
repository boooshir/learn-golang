package repository

import "go-filed-errors/internals/model"

// repository contract
type PostRepository interface {
	CreatePost(post *model.Post) error
}

type postRepository struct{}

func NewPostRepository() PostRepository {
	return &postRepository{}
}

func (r *postRepository) CreatePost(post *model.Post) error {
	return nil
}

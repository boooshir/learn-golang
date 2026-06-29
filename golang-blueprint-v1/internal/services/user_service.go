package services

import (
	"context"
	"errors"
	"fmt"
	"golang-blueprint-v1/internal/models"
	"golang-blueprint-v1/internal/repositories"
)

type UserService interface {
	Create(ctx context.Context, input *models.RegisterRequest) error
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserServiceImpl(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Create(ctx context.Context, input *models.RegisterRequest) error {
	// find email first
	user, err := s.repo.FindUserByEmail(ctx, input.Email)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	if user.Email != "" {
		return errors.New("email already registerd")
	}
	if err := s.repo.Create(ctx, input); err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

package service

import (
	"errors"
	"learn-go-unit-test/entity"
	"learn-go-unit-test/repository"
)

type CategoryService struct {
	Respository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Respository.FindById(id)
	if category == nil {
		return category, errors.New("category Not Found")
	} else {
		return category, nil
	}
}

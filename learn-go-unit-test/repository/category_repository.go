package repository

import "learn-go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

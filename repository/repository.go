package repository

import "github.com/bayusamudra5502/belajar-go-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
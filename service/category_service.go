package service

import (
	"errors"

	"github.com/bayusamudra5502/belajar-go-testing/entity"
	"github.com/bayusamudra5502/belajar-go-testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("category not found")
	}else{
		return category, nil
	}
}
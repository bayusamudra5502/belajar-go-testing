package service_test

import (
	"testing"

	"github.com/bayusamudra5502/belajar-go-testing/entity"
	"github.com/bayusamudra5502/belajar-go-testing/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CategoryRepoMock struct {
	Mock mock.Mock
}

func (repo CategoryRepoMock) FindById(id string) *entity.Category {
	arg := repo.Mock.Called(id)

	if arg.Get(0) == nil {
		return nil
	} else {
		cat := arg.Get(0).(entity.Category)
		return &cat
	}
}

/* Testing */

var mockRepo = &CategoryRepoMock{Mock: mock.Mock{}}
var mockService = service.CategoryService{mockRepo}

func TestCategoryServiceGet(t *testing.T) {
	// Mock Programming
	mockRepo.Mock.On("FindById", "100").Return(nil)
	mockRepo.Mock.On("FindById", "1").Return(entity.Category{
		Id: "1",
		Name: "Belalang",
		Price: 100,
	})

	tableTest := []struct{
		name string
		input string
		expectedValue entity.Category
		isErr bool
	}{
		{"TestId1", "1", entity.Category{
			Id: "1",
			Name: "Belalang",
			Price: 100,
		},false},
		{"TestId2", "100", entity.Category{}, true},
	}

	for _, i := range tableTest {
		t.Run(i.name, func(t *testing.T) {
			result, err := mockService.Get(i.input)

			if(i.isErr){
				assert.Nil(t, result, "result must be nil")
				assert.NotNil(t, err, "error must not be nil")
			} else {
				assert.Nil(t, err, "err must be nil")
				assert.NotNil(t, result, "result must not be nil")
				assert.Equal(t, i.expectedValue, *result)
			}
		})
	}
}
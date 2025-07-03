// service/car_service_test.go
package service

import (
	"p2final/model"
	"p2final/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllAvailable(t *testing.T) {
	mockRepo := new(repository.MockCarRepository)
	service := NewCarService(mockRepo)

	expectedCars := []model.Car{
		{ID: 1, Name: "Avanza", IsAvailable: true},
		{ID: 2, Name: "Brio", IsAvailable: true},
	}

	mockRepo.On("GetAllAvailable").Return(expectedCars, nil)

	result, err := service.GetAllAvailable()

	assert.NoError(t, err)
	assert.Equal(t, expectedCars, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateCar(t *testing.T) {
	mockRepo := new(repository.MockCarRepository)
	service := NewCarService(mockRepo)

	input := model.Car{Name: "Jazz", Code: "JZ01", RentalCost: 300000, Category: "small car"}
	output := model.Car{ID: 1, Name: "Jazz", Code: "JZ01", RentalCost: 300000, Category: "small car"}

	mockRepo.On("Create", input).Return(output, nil)

	result, err := service.CreateCar(input)

	assert.NoError(t, err)
	assert.Equal(t, output, result)
	mockRepo.AssertExpectations(t)
}

func TestGetOwnedCars(t *testing.T) {
	mockRepo := new(repository.MockCarRepository)
	service := NewCarService(mockRepo)

	expectedCars := []model.Car{
		{ID: 1, Name: "Innova", OwnerID: 5},
	}

	mockRepo.On("GetOwnedCars", uint(5)).Return(expectedCars, nil)

	result, err := service.GetOwnedCars(5)

	assert.NoError(t, err)
	assert.Equal(t, expectedCars, result)
	mockRepo.AssertExpectations(t)
}

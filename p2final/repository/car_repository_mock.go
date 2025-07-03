// repository/mock_car_repository.go
package repository

import (
	"p2final/model"

	"github.com/stretchr/testify/mock"
)

type MockCarRepository struct {
	mock.Mock
}

func (m *MockCarRepository) Create(car model.Car) (model.Car, error) {
	args := m.Called(car)
	return args.Get(0).(model.Car), args.Error(1)
}

func (m *MockCarRepository) GetAllAvailable() ([]model.Car, error) {
	args := m.Called()
	return args.Get(0).([]model.Car), args.Error(1)
}

func (m *MockCarRepository) GetByID(id uint) (model.Car, error) {
	args := m.Called(id)
	return args.Get(0).(model.Car), args.Error(1)
}

func (m *MockCarRepository) Update(car model.Car) (model.Car, error) {
	args := m.Called(car)
	return args.Get(0).(model.Car), args.Error(1)
}

func (m *MockCarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCarRepository) GetOwnedCars(userID uint) ([]model.Car, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Car), args.Error(1)
}

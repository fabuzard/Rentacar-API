// repository/mock_booking_repository.go
package repository

import (
	"p2final/model"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockBookingRepository struct {
	mock.Mock
}

func (m *MockBookingRepository) GetCarByID(carID uint) (model.Car, error) {
	args := m.Called(carID)
	return args.Get(0).(model.Car), args.Error(1)
}

func (m *MockBookingRepository) GetUserByID(userID uint) (model.User, error) {
	args := m.Called(userID)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockBookingRepository) UpdateUser(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockBookingRepository) UpdateCar(car model.Car) error {
	args := m.Called(car)
	return args.Error(0)
}

func (m *MockBookingRepository) CreateRental(r model.RentalHistory) (model.RentalHistory, error) {
	args := m.Called(r)
	return args.Get(0).(model.RentalHistory), args.Error(1)
}

func (m *MockBookingRepository) CreateTransaction(tx model.TransactionHistory) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockBookingRepository) ReturnCar(rentalID uint, returnTime time.Time) error {
	args := m.Called(rentalID, returnTime)
	return args.Error(0)
}

func (m *MockBookingRepository) GetRentalByID(rentalID uint) (model.RentalHistory, error) {
	args := m.Called(rentalID)
	return args.Get(0).(model.RentalHistory), args.Error(1)
}

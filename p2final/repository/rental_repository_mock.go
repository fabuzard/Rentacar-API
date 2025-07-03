package repository

import (
	"p2final/model"

	"github.com/stretchr/testify/mock"
)

type MockRentalRepository struct {
	mock.Mock
}

func (m *MockRentalRepository) GetUserRentalHistories(userID uint) ([]model.RentalHistory, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.RentalHistory), args.Error(1)
}

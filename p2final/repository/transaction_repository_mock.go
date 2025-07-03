package repository

import (
	"p2final/model"

	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) GetByUserID(userID uint) ([]model.TransactionHistory, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.TransactionHistory), args.Error(1)
}

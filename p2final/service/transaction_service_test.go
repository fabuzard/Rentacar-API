package service

import (
	"p2final/model"
	"p2final/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserTransactions(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)
	svc := NewTransactionService(mockRepo)

	expectedTransactions := []model.TransactionHistory{
		{SenderID: 1, ReceiverID: 2, Amount: 100000, Description: "Topup"},
		{SenderID: 1, ReceiverID: 3, Amount: 50000, Description: "Payment"},
	}

	// Set up expectation on mock
	mockRepo.On("GetByUserID", uint(1)).Return(expectedTransactions, nil)

	// Call the service
	result, err := svc.GetUserTransactions(1)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, result)
	mockRepo.AssertExpectations(t)
}

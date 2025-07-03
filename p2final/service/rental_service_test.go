package service

import (
	"p2final/model"
	"p2final/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserRentalHistories(t *testing.T) {
	mockRepo := new(repository.MockRentalRepository)
	svc := NewRentalService(mockRepo)

	expectedRentalHistory := []model.RentalHistory{
		{ID: 1, UserID: 2, CarID: 2, Cost: 100000},
		{ID: 2, UserID: 2, CarID: 3, Cost: 100000},
	}
	mockRepo.On("GetUserRentalHistories", uint(1)).Return(expectedRentalHistory, nil)

	result, err := svc.GetUserRentalHistories(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRentalHistory, result)
	mockRepo.AssertExpectations(t)
}

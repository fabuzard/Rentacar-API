// service/user_service_test.go
package service

import (
	"p2final/model"
	"p2final/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopUp_Success(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewUserService(mockRepo)

	user := model.User{ID: 1, Balance: 100000}
	expected := model.User{ID: 1, Balance: 150000}

	mockRepo.On("GetByID", uint(1)).Return(user, nil)
	mockRepo.On("UpdateBalance", uint(1), 150000).Return(expected, nil)

	result, err := service.TopUp(1, 50000)

	assert.NoError(t, err)
	assert.Equal(t, expected.Balance, result.Balance)
	mockRepo.AssertExpectations(t)
}

func TestTopUp_InvalidAmount(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewUserService(mockRepo)

	result, err := service.TopUp(1, 0)

	assert.Error(t, err)
	assert.Equal(t, model.User{}, result)
}

func TestGetByEmail(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewUserService(mockRepo)

	expected := model.User{ID: 1, Email: "john@example.com"}
	mockRepo.On("GetByEmail", "john@example.com").Return(expected, nil)

	result, err := service.GetByEmail("john@example.com")

	assert.NoError(t, err)
	assert.Equal(t, expected.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewUserService(mockRepo)

	expected := model.User{ID: 1, Name: "John"}
	mockRepo.On("GetByID", uint(1)).Return(expected, nil)

	result, err := service.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expected.ID, result.ID)
	mockRepo.AssertExpectations(t)
}

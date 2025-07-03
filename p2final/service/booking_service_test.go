// service/booking_service_test.go
package service

import (
	"p2final/model"
	"p2final/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookCar_Success(t *testing.T) {
	mockRepo := new(repository.MockBookingRepository)
	svc := NewBookingService(mockRepo)

	car := model.Car{ID: 1, RentalCost: 100000, IsAvailable: true, OwnerID: 2}
	renter := model.User{ID: 1, Balance: 200000}
	owner := model.User{ID: 2, Balance: 300000}
	rentalInput := model.RentalHistory{UserID: 1, CarID: 1, Cost: 100000}
	rentalResult := rentalInput
	rentalResult.ID = 10
	transaction := model.TransactionHistory{
		SenderID:    1,
		ReceiverID:  2,
		Amount:      100000,
		Description: "Car rental payment",
		RentalID:    10,
	}

	mockRepo.On("GetCarByID", uint(1)).Return(car, nil)
	mockRepo.On("GetUserByID", uint(1)).Return(renter, nil)
	mockRepo.On("GetUserByID", uint(2)).Return(owner, nil)
	mockRepo.On("UpdateUser", mock.Anything).Return(nil).Twice()
	mockRepo.On("UpdateCar", mock.Anything).Return(nil)
	mockRepo.On("CreateRental", mock.AnythingOfType("model.RentalHistory")).Return(rentalResult, nil)
	mockRepo.On("CreateTransaction", transaction).Return(nil)

	result, err := svc.BookCar(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, rentalResult.ID, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestReturnCar_Success(t *testing.T) {
	mockRepo := new(repository.MockBookingRepository)
	svc := NewBookingService(mockRepo)

	rental := model.RentalHistory{ID: 1, UserID: 1, ReturnAt: nil}
	mockRepo.On("GetRentalByID", uint(1)).Return(rental, nil)
	mockRepo.On("ReturnCar", uint(1), mock.AnythingOfType("time.Time")).Return(nil)

	err := svc.ReturnCar(1, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

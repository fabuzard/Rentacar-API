package service

import (
	"errors"
	"fmt"
	"p2final/model"
	"p2final/repository"
	"time"
)

type BookingService interface {
	BookCar(renterID, carID uint) (model.RentalHistory, error)
	ReturnCar(userID, rentalID uint) error
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(r repository.BookingRepository) BookingService {
	return &bookingService{repo: r}
}
func (s *bookingService) ReturnCar(userID, rentalID uint) error {
	now := time.Now()

	rental, err := s.repo.GetRentalByID(rentalID)
	if err != nil {
		return fmt.Errorf("rental not found")
	}
	if rental.UserID != userID {
		return fmt.Errorf("unauthorized to return this car")
	}
	if rental.ReturnAt != nil {
		return fmt.Errorf("car has already been returned")
	}

	return s.repo.ReturnCar(rentalID, now)
}

func (s *bookingService) BookCar(renterID, carID uint) (model.RentalHistory, error) {
	car, err := s.repo.GetCarByID(carID)
	if err != nil {
		return model.RentalHistory{}, errors.New("car not found")
	}

	if !car.IsAvailable {
		return model.RentalHistory{}, errors.New("car is not available")
	}

	renter, err := s.repo.GetUserByID(renterID)
	if err != nil {
		return model.RentalHistory{}, errors.New("renter not found")
	}

	owner, err := s.repo.GetUserByID(car.OwnerID)
	if err != nil {
		return model.RentalHistory{}, errors.New("car owner not found")
	}

	if renter.ID == car.OwnerID {
		return model.RentalHistory{}, errors.New("you cannot rent your own car")
	}

	if renter.Balance < car.RentalCost {
		return model.RentalHistory{}, errors.New("insufficient balance")
	}

	// Update balances
	renter.Balance -= car.RentalCost
	owner.Balance += car.RentalCost

	err = s.repo.UpdateUser(renter)
	if err != nil {
		return model.RentalHistory{}, errors.New("failed to update renter balance")
	}
	err = s.repo.UpdateUser(owner)
	if err != nil {
		return model.RentalHistory{}, errors.New("failed to update owner balance")
	}

	// Mark car as unavailable
	car.IsAvailable = false
	err = s.repo.UpdateCar(car)
	if err != nil {
		return model.RentalHistory{}, errors.New("failed to update car availability")
	}

	// Create rental record
	rental := model.RentalHistory{
		UserID:   renter.ID,
		CarID:    car.ID,
		Cost:     car.RentalCost,
		RentedAt: time.Now(),
	}

	rental, err = s.repo.CreateRental(rental)
	if err != nil {
		return model.RentalHistory{}, errors.New("failed to create rental history")
	}

	// Create transaction record
	transaction := model.TransactionHistory{
		SenderID:    renter.ID,
		ReceiverID:  owner.ID,
		Amount:      car.RentalCost,
		Description: "Car rental payment",
		RentalID:    rental.ID,
	}
	err = s.repo.CreateTransaction(transaction)
	if err != nil {
		return model.RentalHistory{}, errors.New("failed to create transaction")
	}

	return rental, nil
}

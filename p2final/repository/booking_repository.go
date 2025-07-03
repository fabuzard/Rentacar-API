package repository

import (
	"p2final/model"
	"time"

	"gorm.io/gorm"
)

type BookingRepository interface {
	GetCarByID(carID uint) (model.Car, error)
	GetUserByID(userID uint) (model.User, error)
	UpdateUser(user model.User) error
	UpdateCar(car model.Car) error
	CreateRental(r model.RentalHistory) (model.RentalHistory, error)
	CreateTransaction(t model.TransactionHistory) error
	ReturnCar(rentalID uint, returnTime time.Time) error
	GetRentalByID(rentalID uint) (model.RentalHistory, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}
func (r *bookingRepository) GetRentalByID(rentalID uint) (model.RentalHistory, error) {
	var rental model.RentalHistory
	err := r.db.First(&rental, rentalID).Error
	return rental, err
}
func (r *bookingRepository) GetCarByID(carID uint) (model.Car, error) {
	var car model.Car
	err := r.db.First(&car, carID).Error
	return car, err
}

func (r *bookingRepository) GetUserByID(userID uint) (model.User, error) {
	var user model.User
	err := r.db.First(&user, userID).Error
	return user, err
}

func (r *bookingRepository) UpdateUser(user model.User) error {
	return r.db.Save(&user).Error
}

func (r *bookingRepository) UpdateCar(car model.Car) error {
	return r.db.Save(&car).Error
}

func (r *bookingRepository) CreateRental(rental model.RentalHistory) (model.RentalHistory, error) {
	if err := r.db.Create(&rental).Error; err != nil {
		return model.RentalHistory{}, err
	}

	// Preload Car and User after creation
	if err := r.db.
		Preload("Car").
		Preload("Car.Owner").
		Preload("User").
		First(&rental, rental.ID).Error; err != nil {
		return model.RentalHistory{}, err
	}

	return rental, nil
}

func (r *bookingRepository) CreateTransaction(tx model.TransactionHistory) error {
	return r.db.Create(&tx).Error
}

// return car
func (r *bookingRepository) ReturnCar(rentalID uint, returnTime time.Time) error {
	var rental model.RentalHistory
	if err := r.db.First(&rental, rentalID).Error; err != nil {
		return err
	}

	// Update return time
	rental.ReturnAt = &returnTime
	if err := r.db.Save(&rental).Error; err != nil {
		return err
	}

	// Update car availability
	var car model.Car
	if err := r.db.First(&car, rental.CarID).Error; err != nil {
		return err
	}
	car.IsAvailable = true
	return r.db.Save(&car).Error
}

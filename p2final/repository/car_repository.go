package repository

import (
	"p2final/model"

	"gorm.io/gorm"
)

type CarRepository interface {
	Create(car model.Car) (model.Car, error)
	GetAllAvailable() ([]model.Car, error)
	GetByID(id uint) (model.Car, error)
	Update(car model.Car) (model.Car, error)
	Delete(id uint) error
	GetOwnedCars(userID uint) ([]model.Car, error)
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db}
}

func (r *carRepository) Create(car model.Car) (model.Car, error) {
	err := r.db.Create(&car).Error
	return car, err
}

func (r *carRepository) GetAllAvailable() ([]model.Car, error) {
	var cars []model.Car
	err := r.db.Where("is_available = ?", true).Find(&cars).Error
	return cars, err
}

func (r *carRepository) GetByID(id uint) (model.Car, error) {
	var car model.Car
	err := r.db.First(&car, id).Error
	return car, err
}

func (r *carRepository) Update(car model.Car) (model.Car, error) {
	err := r.db.Save(&car).Error
	return car, err
}

func (r *carRepository) Delete(id uint) error {
	return r.db.Delete(&model.Car{}, id).Error
}

func (r *carRepository) GetOwnedCars(userID uint) ([]model.Car, error) {
	var cars []model.Car
	err := r.db.Where("owner_id = ?", userID).Find(&cars).Error
	return cars, err
}

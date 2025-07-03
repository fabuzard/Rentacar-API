package repository

import (
	"p2final/model"

	"gorm.io/gorm"
)

type RentalRepository interface {
	GetUserRentalHistories(userID uint) ([]model.RentalHistory, error)
}

type rentalRepository struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &rentalRepository{db}
}

func (r *rentalRepository) GetUserRentalHistories(userID uint) ([]model.RentalHistory, error) {
	var histories []model.RentalHistory
	err := r.db.Preload("Car").Where("user_id = ?", userID).Find(&histories).Error
	return histories, err
}

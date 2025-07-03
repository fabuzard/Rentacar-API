package repository

import (
	"p2final/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
	GetByEmail(email string) (model.User, error)
	UpdateBalance(userID uint, newBalance int) (model.User, error)
	GetByID(userID uint) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) GetByID(userID uint) (model.User, error) {
	var user model.User
	err := r.db.First(&user, userID).Error
	return user, err
}

// update balance (top up)

func (r *userRepository) UpdateBalance(userID uint, newBalance int) (model.User, error) {
	var user model.User
	err := r.db.Model(&user).Where("id = ?", userID).Update("balance", newBalance).First(&user).Error
	return user, err
}

package repository

import (
	"p2final/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetByUserID(userID uint) ([]model.TransactionHistory, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetByUserID(userID uint) ([]model.TransactionHistory, error) {
	var txs []model.TransactionHistory
	err := r.db.
		Preload("Receiver").
		Preload("Sender").
		Preload("Rental").
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Order("created_at DESC").
		Find(&txs).Error
	return txs, err
}

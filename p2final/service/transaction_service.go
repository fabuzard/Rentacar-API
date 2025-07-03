package service

import (
	"p2final/model"
	"p2final/repository"
)

type TransactionService interface {
	GetUserTransactions(userID uint) ([]model.TransactionHistory, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return &transactionService{repo: r}
}

func (s *transactionService) GetUserTransactions(userID uint) ([]model.TransactionHistory, error) {
	return s.repo.GetByUserID(userID)
}

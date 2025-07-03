package service

import (
	"p2final/model"
	"p2final/repository"
)

type RentalService interface {
	GetUserRentalHistories(userID uint) ([]model.RentalHistory, error)
}

type rentalService struct {
	repo repository.RentalRepository
}

func NewRentalService(r repository.RentalRepository) RentalService {
	return &rentalService{repo: r}
}

func (s *rentalService) GetUserRentalHistories(userID uint) ([]model.RentalHistory, error) {
	return s.repo.GetUserRentalHistories(userID)
}

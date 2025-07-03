package service

import (
	"errors"
	"p2final/model"
	"p2final/repository"
)

type CarService interface {
	CreateCar(car model.Car) (model.Car, error)
	GetAllAvailable() ([]model.Car, error)
	GetByID(id uint) (model.Car, error)
	UpdateCar(car model.Car) (model.Car, error)
	DeleteCar(id uint) error
	GetOwnedCars(userID uint) ([]model.Car, error)
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(r repository.CarRepository) CarService {
	return &carService{repo: r}
}

func (s *carService) CreateCar(car model.Car) (model.Car, error) {
	if car.Name == "" || car.Code == "" || car.Category == "" || car.RentalCost <= 0 {
		return model.Car{}, errors.New("name, code, category, and rental_cost are required")
	}
	return s.repo.Create(car)
}

func (s *carService) GetAllAvailable() ([]model.Car, error) {
	return s.repo.GetAllAvailable()
}

func (s *carService) GetByID(id uint) (model.Car, error) {
	return s.repo.GetByID(id)
}

func (s *carService) UpdateCar(car model.Car) (model.Car, error) {
	return s.repo.Update(car)
}

func (s *carService) DeleteCar(id uint) error {
	return s.repo.Delete(id)
}

func (s *carService) GetOwnedCars(userID uint) ([]model.Car, error) {
	return s.repo.GetOwnedCars(userID)
}

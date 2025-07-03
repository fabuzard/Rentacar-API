package service

import (
	"fmt"
	"p2final/dto"
	"p2final/helper"
	"p2final/model"
	"p2final/repository"
)

type UserService interface {
	CreateUser(user dto.RegisterRequest) (model.User, error)
	GetByEmail(email string) (model.User, error)
	TopUp(userID uint, amount int) (model.User, error)
	GetByID(userID uint) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetByID(userID uint) (model.User, error) {
	return s.repo.GetByID(userID)
}
func (s *userService) CreateUser(user dto.RegisterRequest) (model.User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		// return model.User{}, errors.New("name, email, password wajib diisi")
		// UBAH: error code + pesan -> pakai fmt + constant
		return model.User{}, fmt.Errorf("%s: name,email,password is required", ErrValidation)
	}
	// Tambah: cek email sudah ada atau belum (pseudocode)
	existing, _ := s.repo.GetByEmail(user.Email)
	if existing.ID != 0 {
		return model.User{}, fmt.Errorf("%s: email already used", ErrEmailAlreadyExist)
	}

	// Email validation via VerifyRight
	isValid, reason, err := helper.IsEmailValid(user.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("email validation error: %v", err)
	}
	if !isValid {
		return model.User{}, fmt.Errorf("%s: email invalid - %s", ErrValidation, reason)
	}

	// Hash password
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("failed Hashing password :%s", err)
	}

	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		Balance:  0,
	}

	return s.repo.Create(newUser)
}

func (s *userService) GetByEmail(email string) (model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *userService) TopUp(userID uint, amount int) (model.User, error) {
	if amount <= 0 {
		return model.User{}, fmt.Errorf("amount must be greater than 0")
	}
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return model.User{}, err
	}
	user.Balance += amount
	return s.repo.UpdateBalance(user.ID, user.Balance)
}

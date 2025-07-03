package dto

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type TopUpRequest struct {
	Amount int `json:"amount" validate:"required,gt=0"`
}

// --- CAR ---
type CreateCarRequest struct {
	Name       string `json:"name" validate:"required"`
	Code       string `json:"code" validate:"required"`
	Category   string `json:"category" validate:"required"`
	RentalCost int    `json:"rental_cost" validate:"required,gt=0"`
}

type UpdateCarRequest struct {
	Name       string `json:"name,omitempty"`
	Code       string `json:"code,omitempty"`
	Category   string `json:"category,omitempty"`
	RentalCost int    `json:"rental_cost,omitempty"`
}

// --- RENTAL ---
type BookCarRequest struct {
	CarID uint `json:"car_id" validate:"required"`
}
type ReturnCarRequest struct {
	RentalID uint `json:"rental_id" validate:"required"`
}

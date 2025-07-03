package dto

import "time"

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TopupResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type CreateCarResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Category string `json:"category"`
}

type GetCarResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Category    string `json:"category"`
	RentalCost  int    `json:"rental_cost"`
	IsAvailable bool   `json:"is_available"`
}

type DeletedCarResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type RentalHistoryResponse struct {
	ID       uint       `json:"id"`
	CarName  string     `json:"car_name"`
	Cost     int        `json:"cost"`
	RentedAt time.Time  `json:"rented_at"`
	ReturnAt *time.Time `json:"return_at,omitempty"`
}

type BookingResponse struct {
	RentalID    uint      `json:"rental_id"`
	CarID       uint      `json:"car_id"`
	CarName     string    `json:"car_name"`
	Category    string    `json:"category"`
	Cost        int       `json:"cost"`
	RentedAt    time.Time `json:"rented_at"`
	UserBalance int       `json:"your_balance"`
}

type TransactionUserInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TransactionResponse struct {
	ID          uint                `json:"id"`
	Sender      TransactionUserInfo `json:"sender"`
	Receiver    TransactionUserInfo `json:"receiver"`
	Amount      int                 `json:"amount"`
	Description string              `json:"description"`
	RentalID    uint                `json:"rental_id"`
	CreatedAt   string              `json:"created_at"`
}

type MeResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Balance int    `json:"balance"`
}

package model

import (
	"gorm.io/gorm"
)

type TransactionHistory struct {
	gorm.Model
	SenderID    uint
	Sender      User
	ReceiverID  uint
	Receiver    User
	Amount      int `gorm:"not null"`
	Description string
	RentalID    uint
	Rental      RentalHistory `gorm:"foreignKey:RentalID"`
}

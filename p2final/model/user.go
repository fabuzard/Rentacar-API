package model

import "time"

type User struct {
	ID               uint                 `gorm:"primaryKey"`
	Name             string               `gorm:"type:varchar(100);not null"`
	Email            string               `gorm:"type:varchar(100);not null;unique"`
	Password         string               `gorm:"not null"`
	Balance          int                  `gorm:"not null;default:0"`
	CreatedAt        time.Time            `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time            `gorm:"default:CURRENT_TIMESTAMP"`
	RentalHistories  []RentalHistory      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	TransactionsSent []TransactionHistory `gorm:"foreignKey:SenderID"`
	TransactionsRecv []TransactionHistory `gorm:"foreignKey:ReceiverID"`
}

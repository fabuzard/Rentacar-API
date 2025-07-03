package model

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Code        string `gorm:"type:varchar(50);unique;not null"`
	Category    string `gorm:"type:varchar(50);not null"`
	RentalCost  int    `gorm:"not null"`
	IsAvailable bool   `gorm:"not null;default:true"`

	OwnerID uint
	Owner   User

	RentalHistories []RentalHistory `gorm:"foreignKey:CarID;constraint:OnDelete:CASCADE"`
}

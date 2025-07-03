package model

import "time"

type RentalHistory struct {
	ID       uint      `gorm:"primaryKey"`
	UserID   uint      `gorm:"not null"`
	CarID    uint      `gorm:"not null"`
	Cost     int       `gorm:"not null"`
	RentedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ReturnAt *time.Time

	User User `gorm:"foreignKey:UserID"`
	Car  Car  `gorm:"foreignKey:CarID"`
}

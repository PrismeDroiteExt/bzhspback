package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	OrderId       uint   `gorm:"not null"`
	UserId        uint   `gorm:"not null"`
	Status        string `gorm:"not null"`
	PaymentMethod string `gorm:"not null"`
	TransationId  string `gorm:"not null"`
}

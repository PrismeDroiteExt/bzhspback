package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model

	CartId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	Quantity  uint `gorm:"not null"`
}

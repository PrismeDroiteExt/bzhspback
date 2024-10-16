package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	CartId uint   `gorm:"not null"`
	Status string `gorm:"not null"`
}

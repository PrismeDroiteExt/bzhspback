package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	// Pictures    []Picture
	Price float64 `gorm:"not null"`
}

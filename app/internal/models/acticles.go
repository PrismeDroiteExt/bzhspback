package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	TagID       int    `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	// Pictures    []Picture
	Price float64 `gorm:"not null"`
}

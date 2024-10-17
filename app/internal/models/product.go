package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Title       string   `gorm:"not null"`
	Description string   `gorm:"not null"`
	CategoryID  uint     `gorm:"not null"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Price       float64  `gorm:"not null"`
	Discount    *float64 `gorm:"default:null"`
	Colors      string   `gorm:"not null"`
	Sizes       string   `gorm:"not null"`
	PictureUrl  string   `gorm:"not null"`
}

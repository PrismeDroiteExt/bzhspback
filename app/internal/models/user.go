package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	TagID        int    `gorm:"primaryKey;autoIncrement"`
	Mail         string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	RegisterDate string `gorm:"not null"`
	Name         string `gorm:"not null"`
	LastName     string `gorm:"not null"`
}

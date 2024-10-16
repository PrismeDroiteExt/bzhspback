package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	UserId      User `gorm:"references:ID"`
	IsProcessed bool `gorm:"not null"`
}

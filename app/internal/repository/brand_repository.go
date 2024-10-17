package repository

import (
	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/gorm"
)

// BrandRepository struct represents a repository for managing brand data.
type BrandRepository struct {
	DB *gorm.DB
}

// Constructor for the BrandRepository struct.
func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{DB: db}
}

// GetAllBrands retrieves all brands from the database.
func (r *BrandRepository) GetAllBrands() ([]models.Brand, error) {
	var brands []models.Brand
	result := r.DB.Find(&brands)
	return brands, result.Error
}

// Get Brand by ID
func (r *BrandRepository) GetBrandByID(id uint) (models.Brand, error) {
	var brand models.Brand
	result := r.DB.First(&brand, id)
	return brand, result.Error
}

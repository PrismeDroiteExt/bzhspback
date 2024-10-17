package repository

import (
	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/gorm"
)

// ProductRepository struct represents a repository for managing product data.
type ProductRepository struct {
	DB *gorm.DB
}

// Constructor for the ProductRepository struct.
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// Get Product by ID
func (r *ProductRepository) GetProductByID(id uint) (models.Product, error) {
	var product models.Product

	result := r.DB.Preload("Category").First(&product, id)

	return product, result.Error
}

// Get all products for a category
func (r *ProductRepository) GetProductsByCategoryID(categoryID uint) ([]models.Product, error) {
	var products []models.Product

	query := `
		WITH RECURSIVE category_tree AS (
			SELECT id FROM categories WHERE id = ?
			UNION ALL
			SELECT c.id FROM categories c
			INNER JOIN category_tree ct ON c.sub_category_id = ct.id
		)
		SELECT p.* FROM products p
		INNER JOIN category_tree ct ON p.category_id = ct.id
	`

	result := r.DB.Raw(query, categoryID).Preload("Brand").Preload("Category").Find(&products)

	return products, result.Error
}

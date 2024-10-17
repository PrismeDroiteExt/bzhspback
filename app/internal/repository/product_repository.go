package repository

import (
	"fmt"

	"bzhspback.fr/breizhsport/internal/dto"
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
func (r *ProductRepository) GetProductsByCategoryID(categoryID uint, filters []dto.Filter) ([]models.Product, error) {
	var products []models.Product

	query := `
		WITH RECURSIVE category_tree AS (
			SELECT id FROM categories WHERE id = ?
			UNION ALL
			SELECT c.id FROM categories c
			INNER JOIN category_tree ct ON c.sub_category_id = ct.id
		)
		SELECT DISTINCT p.* FROM products p
		INNER JOIN category_tree ct ON p.category_id = ct.id
		WHERE 1=1
	`
	args := []interface{}{categoryID}

	for _, filter := range filters {
		switch filter.Field {
		case "brand_id":
			query += " AND p.brand_id IN (?)"
			args = append(args, filter.Value)
		case "min_price":
			query += " AND p.price >= ?"
			args = append(args, filter.Value)
		case "max_price":
			query += " AND p.price <= ?"
			args = append(args, filter.Value)
		case "colors", "sizes":
			query += fmt.Sprintf(" AND string_to_array(p.%s, ',') && string_to_array(?, ',')", filter.Field)
			if str, ok := filter.Value.(string); ok {
				args = append(args, str)
			} else {
				return nil, fmt.Errorf("invalid type for %s filter: expected string", filter.Field)
			}
		}
	}

	// TODO: implement the pagination et the limit

	result := r.DB.Raw(query, args...).Preload("Brand").Preload("Category").Find(&products)

	return products, result.Error
}

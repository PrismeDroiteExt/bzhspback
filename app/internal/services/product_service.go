package services

import (
	"bzhspback.fr/breizhsport/internal/dto"
	"bzhspback.fr/breizhsport/internal/repository"
	"github.com/jinzhu/copier"
)

// ProductService struct represents a service for managing product data.
type ProductService struct {
	repo *repository.ProductRepository
}

// Constructor
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Get Product by ID
func (s *ProductService) GetProductByID(id uint) (dto.ProductResponse, error) {
	product, err := s.repo.GetProductByID(id)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	var productDTO dto.ProductResponse

	err = copier.Copy(&productDTO, &product)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	return productDTO, nil
}

// Get all products for a category
func (s *ProductService) GetProductsByCategoryID(categoryID uint) ([]dto.ProductResponse, error) {
	products, err := s.repo.GetProductsByCategoryID(categoryID)

	if err != nil {
		return []dto.ProductResponse{}, err
	}

	var productDTOs []dto.ProductResponse

	err = copier.Copy(&productDTOs, &products)

	if err != nil {
		return []dto.ProductResponse{}, err
	}

	return productDTOs, nil
}
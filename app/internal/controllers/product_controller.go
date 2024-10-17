package controllers

import (
	"net/http"
	"strconv"

	"bzhspback.fr/breizhsport/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *services.ProductService
}

// Constructor
func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// GetProductByID handles the GET request to retrieve a product by its ID.
// It returns a JSON response with the product details or an error message.
func (c *ProductController) GetProductByID(ctx *gin.Context) {
	// get id parameter
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := c.service.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// GetProductsByCategoryID handles the GET request to retrieve all products for a category by its ID.
// It returns a JSON response with the products details or an error message.
func (c *ProductController) GetProductsByCategoryID(ctx *gin.Context) {
	// get id parameter
	categoryID, err := strconv.Atoi(ctx.Param("category_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	products, err := c.service.GetProductsByCategoryID(uint(categoryID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

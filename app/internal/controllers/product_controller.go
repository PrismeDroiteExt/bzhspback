package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"bzhspback.fr/breizhsport/internal/dto"
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
	categoryID, err := strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Parse query parameters
	var filterRequest dto.FilterRequest
	filterRequest.Filters = make([]dto.Filter, 0)

	// Handle brand
	if brandIDsStr := ctx.Query("brand"); brandIDsStr != "" {
		brandIDs := strings.Split(brandIDsStr, ",")
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "brand_id", Value: brandIDs})
	}

	// Handle min_price
	if minPrice := ctx.Query("min_price"); minPrice != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "min_price", Value: minPrice})
	}

	// Handle max_price
	if maxPrice := ctx.Query("max_price"); maxPrice != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "max_price", Value: maxPrice})
	}

	// Handle colors
	if colors := ctx.Query("colors"); colors != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "colors", Value: colors})
	}

	// Handle sizes
	if sizes := ctx.Query("sizes"); sizes != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "sizes", Value: sizes})
	}

	products, err := c.service.GetProductsByCategoryID(uint(categoryID), filterRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

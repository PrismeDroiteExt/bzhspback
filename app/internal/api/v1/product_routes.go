package api

import (
	"bzhspback.fr/breizhsport/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Parameters:
//   - router: A pointer to the gin.RouterGroup where the product routes will be added.
//   - controller: A pointer to the ProductController that handles brand-related operations.
func SetupProductRoutes(router *gin.RouterGroup, controller *controllers.ProductController) {
	productRoutes := router.Group("/products") // Group all product routes under /brands
	{
		// GET /brands - Retrieves all brands
		productRoutes.GET("/:id", controller.GetProductByID)
		productRoutes.GET("/category/:category_id", controller.GetProductsByCategoryID)

		// TODO: Add more brand-related routes here as needed
		// For example:
		// brandRoutes.GET("/:id", controller.GetBrandByID)
		// brandRoutes.POST("", controller.CreateBrand)
		// brandRoutes.PUT("/:id", controller.UpdateBrand)
		// brandRoutes.DELETE("/:id", controller.DeleteBrand)
	}
}

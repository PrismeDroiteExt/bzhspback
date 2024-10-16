package api

import (
	"bzhspback.fr/breizhsport/internal/controllers"
	"bzhspback.fr/breizhsport/internal/repository"
	"bzhspback.fr/breizhsport/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	apiv1 := r.Group("/api/v1")

	initBrandRoutes(apiv1, db)

	// TODO: Add other route setups here (e.g., products, users)
}

func initBrandRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewBrandRepository(db)
	service := services.NewBrandService(repo)
	controller := controllers.NewBrandController(service)
	SetupBrandRoutes(router, controller)
}

package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"users": "To be implemented",
			})
		})
	}
}

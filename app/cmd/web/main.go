package main

import (
	"bzhspback.fr/breizhsport/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run()
}

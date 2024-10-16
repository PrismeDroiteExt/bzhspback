package main

import (
	"bzhspback.fr/breizhsport/internal/api"
	"bzhspback.fr/breizhsport/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDB()
	api.InitRoutes(r)

	r.Run()
}

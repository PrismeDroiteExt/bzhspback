package database

import (
	"fmt"
	"log"
	"os"

	"bzhspback.fr/breizhsport/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Attempting to connect to database")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		fmt.Println("Connected to database")
	}

	// Auto Migrate
	err = DB.AutoMigrate(&models.User{}, &models.Article{})
	if err != nil {
		log.Fatal("Failed to auto migrate:", err)
	} else {
		fmt.Println("Auto migrated database")
	}
}

func GetDB() *gorm.DB {
	return DB
}

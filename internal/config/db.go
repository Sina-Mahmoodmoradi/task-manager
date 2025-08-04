package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Sina-Mahmoodmoradi/task-manager/internal/repository/gorm/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ failed to load .env file: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ failed to connect to database: %v", err)
    }

    DB = db

    err = db.AutoMigrate(&models.User{}, &models.Task{})
    if err != nil {
        log.Fatalf("❌ migration failed: %v", err)
    }

    log.Println("✅ Connected to PostgreSQL")
}

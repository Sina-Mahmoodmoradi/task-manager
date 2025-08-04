package config

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
	"github.com/joho/godotenv"
    "gorm.io/gorm"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/repository/gorm/models"
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

    err = db.AutoMigrate(&models.UserModel{}, &models.TaskModel{})
    if err != nil {
        log.Fatalf("❌ migration failed: %v", err)
    }

    log.Println("✅ Connected to PostgreSQL")
}

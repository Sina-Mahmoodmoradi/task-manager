package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Sina-Mahmoodmoradi/task-manager/internal/repository/gorm/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {
	_ = godotenv.Load()

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
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	log.Println("✅ Connected to PostgreSQL")

	return &Database{DB: db}, nil
}

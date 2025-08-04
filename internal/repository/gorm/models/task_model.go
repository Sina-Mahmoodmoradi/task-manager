package models

import (
	"gorm.io/gorm"
)

type Task struct {
    gorm.Model
    Title  string `gorm:"not null"`
    Done   bool
    UserID uint
}

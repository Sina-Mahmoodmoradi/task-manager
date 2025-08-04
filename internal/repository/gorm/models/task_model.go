package models

import (
    "gorm.io/gorm"
)

type TaskModel struct {
    gorm.Model
    Title  string `gorm:"not null"`
    Done   bool
    UserID uint
}

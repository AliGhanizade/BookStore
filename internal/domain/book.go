package domain

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:100,unique"`
	Author    string `gorm:"size:100"`
	Status    string
	Year      uint
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

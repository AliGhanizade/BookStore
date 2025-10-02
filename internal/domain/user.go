package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           ` gorm:"primaryKey"`
	Name       string         ` gorm:"size:100"`
	Phone      string         `gorm:"uniqueIndex"`
	Password   string         
	Permission string        
	CreatedAt  time.Time     
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

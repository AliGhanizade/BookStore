package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name" gorm:"size:100"`
	Phone      string         `json:"phone" gorm:"uniqueIndex"`
	Password   string         `json:"password"`
	Permission string         `json:"permission"`
	CreatedAt  time.Time     
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

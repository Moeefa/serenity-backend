package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string				 `json:"name"`
	Email       string         `gorm:"uniqueIndex" json:"email"`
	Password    string         `json:"password"`
	Phone       string         `json:"phone"`
	LastActive  time.Time      `json:"last_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}
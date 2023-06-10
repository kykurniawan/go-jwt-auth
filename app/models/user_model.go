package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

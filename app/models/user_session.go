package models

import "time"

type UserSession struct {
	ID           uint      `gorm:"primary_key;auto_increment"`
	UserID       uint      `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	ExpiredAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

package database

import (
	"github.com/kykurniawan/go-jwt-auth/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func CreateConnection() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to create database connection")
	}

	db.AutoMigrate(
		&models.User{},
		&models.UserSession{},
	)

	Connection = db
}

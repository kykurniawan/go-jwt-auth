package database

import (
	"github.com/kykurniawan/go-jwt-auth/app/models"
	"github.com/kykurniawan/go-jwt-auth/configs"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func CreateConnection() {
	var dialector gorm.Dialector

	switch configs.Database().Driver {
	case "mysql":
		dialector = mysql.Open(
			configs.Database().Username +
				":" + configs.Database().Password +
				"@tcp(" + configs.Database().Host +
				":" + configs.Database().Port + ")/" +
				configs.Database().Database +
				"?charset=" + configs.Database().Charset +
				"&parseTime=True&loc=Local")
	case "sqlite":
		dialector = sqlite.Open(configs.Database().Database)
	default:
		panic("invalid database driver")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		panic("failed to create database connection")
	}

	db.AutoMigrate(
		&models.User{},
		&models.UserSession{},
	)

	Connection = db
}

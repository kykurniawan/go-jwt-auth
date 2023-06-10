package configs

import (
	"github.com/kykurniawan/go-jwt-auth/helpers"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
}

func Database() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   helpers.EnvString("DB_DRIVER", "mysql"),
		Host:     helpers.EnvString("DB_HOST", "localhost"),
		Port:     helpers.EnvString("DB_PORT", "3306"),
		Username: helpers.EnvString("DB_USERNAME", "root"),
		Password: helpers.EnvString("DB_PASSWORD", "root"),
		Database: helpers.EnvString("DB_DATABASE", "app"),
		Charset:  helpers.EnvString("DB_CHARSET", "utf8mb4"),
	}
}

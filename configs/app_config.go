package configs

import (
	"strings"

	"github.com/kykurniawan/go-jwt-auth/helpers"
)

type AppConfig struct {
	Name           string
	Mode           string
	Port           string
	Host           string
	TrustedProxies []string
}

func App() *AppConfig {
	trustedProxies := helpers.EnvString("APP_TRUSTED_PROXIES", "127.0.0.1")
	trustedProxiesSlice := strings.Split(trustedProxies, ",")

	return &AppConfig{
		Name:           helpers.EnvString("APP_NAME", "Go JWT Auth"),
		Mode:           helpers.EnvString("APP_MODE", "debug"),
		Port:           helpers.EnvString("APP_PORT", "8080"),
		Host:           helpers.EnvString("APP_HOST", "localhost"),
		TrustedProxies: trustedProxiesSlice,
	}
}

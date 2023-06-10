package configs

import "github.com/kykurniawan/go-jwt-auth/helpers"

type JWTConfig struct {
	SecretKey             string
	AccessTokenExpiresIn  int
	RefreshTokenExpiresIn int
}

func JWT() *JWTConfig {
	return &JWTConfig{
		SecretKey:             helpers.EnvString("JWT_SECRET", "secret"),
		AccessTokenExpiresIn:  helpers.EnvInt("JWT_ACCESS_TOKEN_EXPIRES_IN", 300),
		RefreshTokenExpiresIn: helpers.EnvInt("JWT_REFRESH_TOKEN_EXPIRES_IN", 1800),
	}
}

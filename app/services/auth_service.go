package services

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kykurniawan/go-jwt-auth/app/models"
	"github.com/kykurniawan/go-jwt-auth/app/repositories"
)

type AuthResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthService struct {
	userRepository        *repositories.UserRepository
	userSessionRepository *repositories.UserSessionRepository
}

func NewAuthService(
	userRepository *repositories.UserRepository,
	userSessionRepository *repositories.UserSessionRepository,
) *AuthService {
	return &AuthService{
		userRepository,
		userSessionRepository,
	}
}

func (service *AuthService) Attempt(email string, password string) (*AuthResult, error) {
	user, err := service.userRepository.FindByEmail(email)

	if err != nil {
		return nil, errors.New("email or password is wrong")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("email or password is wrong")
	}

	accessTokenExpiresIn, _ := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRES_IN"), 10, 64)
	accessToken, err := service.GenerateToken(user, "access_token", accessTokenExpiresIn)

	if err != nil {
		return nil, err
	}

	refreshTokenExpiresIn, _ := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRES_IN"), 10, 64)
	refreshToken, err := service.GenerateToken(user, "refresh_token", refreshTokenExpiresIn)

	if err != nil {
		return nil, err
	}

	fmt.Println(user.ID)

	var userSession models.UserSession

	userSession.UserID = user.ID
	userSession.RefreshToken = refreshToken
	userSession.ExpiredAt = time.Now().Add(time.Duration(refreshTokenExpiresIn) * time.Second)

	err = service.userSessionRepository.Create(&userSession)

	if err != nil {
		return nil, err
	}

	return &AuthResult{
		accessToken,
		refreshToken,
	}, nil
}

func (service *AuthService) Refresh(refreshToken string) (*AuthResult, error) {
	userSession, err := service.userSessionRepository.FindByRefreshToken(refreshToken)

	if err != nil {
		return nil, errors.New("refresh token is invalid")
	}

	if userSession.ExpiredAt.Before(time.Now()) {
		return nil, errors.New("refresh token is expired")
	}

	user, err := service.userRepository.FindById(userSession.UserID)

	if err != nil {
		return nil, errors.New("refresh token is invalid")
	}

	accessTokenExpiresIn, _ := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRES_IN"), 10, 64)
	accessToken, err := service.GenerateToken(user, "access_token", accessTokenExpiresIn)

	if err != nil {
		return nil, err
	}

	refreshTokenExpiresIn, _ := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRES_IN"), 10, 64)
	newRefreshToken, err := service.GenerateToken(user, "refresh_token", refreshTokenExpiresIn)

	if err != nil {
		return nil, err
	}

	userSession.RefreshToken = newRefreshToken
	userSession.ExpiredAt = time.Now().Add(time.Duration(refreshTokenExpiresIn) * time.Second)

	err = service.userSessionRepository.Update(userSession)

	if err != nil {
		return nil, err
	}

	return &AuthResult{
		accessToken,
		newRefreshToken,
	}, nil
}

func (service *AuthService) Logout(refreshToken string) error {
	userSession, err := service.userSessionRepository.FindByRefreshToken(refreshToken)

	if err != nil {
		return errors.New("refresh token is invalid")
	}

	err = service.userSessionRepository.Delete(userSession)

	if err != nil {
		return err
	}

	return nil
}

func (service *AuthService) LogoutAll(refreshToken string) error {
	userSession, err := service.userSessionRepository.FindByRefreshToken(refreshToken)

	if err != nil {
		return errors.New("refresh token is invalid")
	}

	err = service.userSessionRepository.DeleteAllByUserId(userSession.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (service *AuthService) GenerateToken(user *models.User, tokenType string, expiresIn int64) (string, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
		err error
	)

	key = []byte(os.Getenv("JWT_SECRET"))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
		"iat": time.Now().Unix(),
		"typ": tokenType,
	})
	s, err = t.SignedString(key)

	return s, err
}

func (service *AuthService) ValidateToken(token string) (*jwt.Token, error) {
	var (
		key []byte
		t   *jwt.Token
		err error
	)

	key = []byte(os.Getenv("JWT_SECRET"))
	t, err = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("token is invalid")
	}

	return t, nil
}

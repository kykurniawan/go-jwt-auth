package repositories

import (
	"github.com/kykurniawan/go-jwt-auth/app/models"
	"gorm.io/gorm"
)

type UserSessionRepository struct {
	db *gorm.DB
}

func NewUserSessionRepository(db *gorm.DB) *UserSessionRepository {
	return &UserSessionRepository{db}
}

func (r *UserSessionRepository) FindByRefreshToken(refreshToken string) (*models.UserSession, error) {
	var userSession models.UserSession

	err := r.db.Where("refresh_token = ?", refreshToken).First(&userSession).Error

	return &userSession, err
}

func (r *UserSessionRepository) FindById(id uint) (*models.UserSession, error) {
	var userSession models.UserSession

	err := r.db.First(&userSession, id).Error

	return &userSession, err
}

func (r *UserSessionRepository) FindByUserId(userId uint) ([]*models.UserSession, error) {
	var userSessions []*models.UserSession

	err := r.db.Where("user_id = ?", userId).Find(&userSessions).Error

	return userSessions, err
}

func (r *UserSessionRepository) Create(userSession *models.UserSession) error {
	err := r.db.Create(&userSession).Error

	return err
}

func (r *UserSessionRepository) Delete(userSession *models.UserSession) error {
	err := r.db.Delete(&userSession).Error

	return err
}

func (r *UserSessionRepository) DeleteAllByUserId(userId uint) error {
	err := r.db.Where("user_id = ?", userId).Delete(&models.UserSession{}).Error

	return err
}

func (r *UserSessionRepository) Update(userSession *models.UserSession) error {
	err := r.db.Save(&userSession).Error

	return err
}
package repositories

import (
	"github.com/kykurniawan/go-jwt-auth/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *UserRepository) FindById(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error

	return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *UserRepository) Create(user *models.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *UserRepository) Update(user *models.User) error {
	err := r.db.Save(&user).Error

	return err
}

func (r *UserRepository) Delete(user *models.User) error {
	err := r.db.Delete(&user).Error

	return err
}

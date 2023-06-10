package services

import (
	"github.com/kykurniawan/go-jwt-auth/app/models"
	"github.com/kykurniawan/go-jwt-auth/app/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository,
	}
}

func (service *UserService) FindAll() ([]*models.User, error) {
	return service.userRepository.FindAll()
}

func (service *UserService) FindByID(id uint) (*models.User, error) {
	return service.userRepository.FindById(id)
}

func (service *UserService) FindByEmail(email string) (*models.User, error) {
	return service.userRepository.FindByEmail(email)
}

func (service *UserService) Create(user *models.User) (*models.User, error) {
	err := service.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) Update(user *models.User) (*models.User, error) {
	err := service.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) Delete(user *models.User) error {
	return service.userRepository.Delete(user)
}

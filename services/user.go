package services

import (
	"restapi/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (s *AuthService) CreateUser(user *models.User) error {

	hashpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashpassword)

	err = s.db.Create(user).Error

	if err != nil {
		return err
	}

	return nil

}

func (s *AuthService) LoginUser(username string, password string) (bool, error) {
	var user models.User

	err := s.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil

}

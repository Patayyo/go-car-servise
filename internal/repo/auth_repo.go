package repo

import (
	"car-service/internal/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	SignUp(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{DB: db}
}

func (r *authRepository) SignUp(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *authRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

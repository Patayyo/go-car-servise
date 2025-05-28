package service

import (
	"car-service/internal/model"
	"car-service/internal/repo"
	"car-service/pkg/utils"
	"errors"

	"gorm.io/gorm"
)

type AuthService struct {
	repo repo.AuthRepository
}

func NewAuthService(r repo.AuthRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) SignUp(input *model.User) error {
	_, err := s.repo.GetByEmail(input.Email)

	if err == nil {
		return errors.New("user already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	input.Password = hashedPassword

	if err := s.repo.SignUp(input); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GetByEmail(email string) (*model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if ok := utils.CheckPasswordHash(password, user.Password); !ok {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

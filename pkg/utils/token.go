package utils

import (
	"car-service/internal/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID uint) (string, error) {
	exp := time.Now().Add(time.Hour * 2)
	sub := model.User{ID: userID}
	claims := jwt.MapClaims{
		"user_id": sub.ID,
		"exp":     exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

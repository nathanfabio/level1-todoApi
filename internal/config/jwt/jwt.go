package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET_KEY = "JWT_SECRET_KEY"

func GenerateJWT(id int, email string) (string, error) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(username string) (string, error) {
	// Membuat token JWT dengan payload yang sesuai
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token berlaku selama 1 jam
	})

	// Menandatangani token dengan secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret_key := os.Getenv("JWT_KEY")

	return token.SignedString([]byte(secret_key))
}

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ParseToken(tokenStr string) (int, string, error) {
	secretKey := os.Getenv("JWT_KEY")

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, "", fmt.Errorf("Invalid JWT token")
	}

	return claims.UserID, claims.Username, nil
}

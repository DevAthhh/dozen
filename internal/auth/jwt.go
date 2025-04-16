package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, email string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "x-vibe",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string) (*CustomClaims, error) {
	tkn, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected tkn")
		}
		return os.Getenv("SECRET_KEY"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tkn.Claims.(*CustomClaims); ok && tkn.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

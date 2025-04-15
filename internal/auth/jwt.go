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
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 12)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "doZen",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
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

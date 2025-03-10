package security

import (
	"ECCO2K/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID uint   `json:"UserID"`
	Name   string `json:"Name"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, name string) (string, error) {
	cfg := config.LoadEnv()
	secret := []byte(cfg.JWT_SECRET)

	claims := Claims{
		UserID: userID,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	cfg := config.LoadEnv()
	secret := []byte(cfg.JWT_SECRET)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

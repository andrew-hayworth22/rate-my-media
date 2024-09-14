package core

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTFields struct {
	userId int
}

type JWTClaims struct {
	jwt.RegisteredClaims
	JWTFields
}

func GenerateJWT(cfg Config, jwtFields JWTFields) (string, error) {
	expirationDuration := time.Hour
	claims := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rate-my-media",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationDuration)),
		},
		JWTFields: jwtFields,
	}

	tok, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.JwtSecret))
	if err != nil {
		return "", err
	}

	return tok, nil
}

func DecodeJWT(cfg Config, tok string) (*jwt.Token, error) {
	token, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

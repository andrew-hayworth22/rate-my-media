package app

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (a *App) GenerateJWT(user AppUser) (string, error) {
	expirationDuration := time.Hour

	tok, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "rate-my-media",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expirationDuration)),
	}).SignedString([]byte(a.jwtSecret))

	if err != nil {
		return "", err
	}

	return tok, nil
}

func (a *App) DecodeJWT(tok string) (*jwt.Token, error) {
	token, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

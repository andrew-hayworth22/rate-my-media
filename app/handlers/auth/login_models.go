package auth

import (
	"context"
	"net/mail"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginRequest) Valid(ctx context.Context) (problems map[string]string) {
	if _, err := mail.ParseAddress(req.Email); err != nil {
		problems["email"] = "Please enter a valid email address"
	}

	if len(req.Password) == 0 {
		problems["password"] = "Please enter a password"
	}

	return
}

type LoginResponse struct {
	Token string `json:"token"`
}

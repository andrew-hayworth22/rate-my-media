package auth

import (
	"context"
	"net/mail"
	"strings"
)

type PostUserRequest struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	DisplayName          string `json:"display_name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (req PostUserRequest) Valid(ctx context.Context) (problems map[string]string) {
	if _, err := mail.ParseAddress(req.Email); err != nil {
		problems["email"] = "Please enter a valid email address"
	}

	trimmedName := strings.TrimSpace(req.Name)
	if len(trimmedName) == 0 {
		problems["name"] = "Please enter a valid name"
	}

	trimmedDisplayName := strings.TrimSpace(req.DisplayName)
	if len(trimmedDisplayName) == 0 {
		problems["display_name"] = "Please enter a valid display name"
	}

	if len(req.Password) < 8 {
		problems["password"] = "Your password must be at least 8 characters long"
	}

	if req.Password != req.PasswordConfirmation {
		problems["password_confirmation"] = "Your password confirmation does not match your password"
	}

	return
}

type AppUser struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

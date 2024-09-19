package auth

import "context"

type PostUserRequest struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	DisplayName          string `json:"display_name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (req PostUserRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = map[string]string{}
	return
}

type AppUser struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

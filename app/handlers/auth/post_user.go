package auth

import (
	"context"
	"net/http"
	"net/mail"
	"strings"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
)

type PostUserRequest struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	DisplayName          string `json:"display_name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (req PostUserRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = map[string]string{}

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

func HandlePostUser(authStore auth.Store) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var req PostUserRequest
			req, problems, err := core.DecodeValid(r, req)
			if err != nil {
				core.EncodeInternalError(w)
				return
			}
			if len(problems) > 0 {
				core.EncodeValidationError(w, problems)
				return
			}

			existingUser, err := authStore.GetUserByEmail(r.Context(), req.Email)
			if err != nil {
				core.EncodeInternalError(w)
				return
			}

			if existingUser.Id != 0 {
				core.EncodeValidationError(w, map[string]string{
					"email": "This email already has an account associated with it",
				})
				return
			}

			dbUser, err := authStore.StoreUser(r.Context(), auth.DbStoreUserRequest{
				Email:       req.Email,
				Name:        req.Name,
				DisplayName: req.DisplayName,
				Password:    req.Password,
			})
			if err != nil {
				core.EncodeInternalError(w)
				return
			}

			appUsr := AppUser{
				Id:          dbUser.Id,
				Email:       dbUser.Email,
				Name:        dbUser.Name,
				DisplayName: dbUser.DisplayName,
			}
			core.Encode(w, http.StatusOK, appUsr)
		})
}

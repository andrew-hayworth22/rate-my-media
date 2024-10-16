package auth

import (
	"context"
	"net/http"
	"net/mail"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = map[string]string{}

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

func HandleLogin(cfg core.Config, authStore auth.Store) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var req LoginRequest
			req, problems, err := core.DecodeValid(r, req)
			if err != nil {
				core.EncodeInternalError(w)
				return
			}
			if len(problems) > 0 {
				core.EncodeValidationError(w, problems)
				return
			}

			dbUser, err := authStore.GetUserByEmail(r.Context(), req.Email)
			if err != nil {
				core.EncodeInternalError(w)
				return
			}

			if dbUser.Id == 0 {
				core.EncodeUnauthorized(w)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(req.Password))
			if err != nil {
				core.EncodeUnauthorized(w)
				return
			}

			fields := core.JWTFields{
				UserId: dbUser.Id,
			}
			tok, err := core.GenerateJWT(cfg, fields)
			if err != nil {
				core.EncodeInternalError(w)
			}

			core.Encode(w, http.StatusOK, LoginResponse{Token: tok})
		},
	)
}

package auth

import (
	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

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

package auth

import (
	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"net/http"
)

type AppUser struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func HandlePostUser(authStore auth.Store, cfg core.Config) http.Handler {
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

			dbUser, err := authStore.StoreUser(r.Context(), auth.DbStoreUserRequest{})
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

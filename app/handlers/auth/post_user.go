package auth

import (
	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"net/http"
)

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

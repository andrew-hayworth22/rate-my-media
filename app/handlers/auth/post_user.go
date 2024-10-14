package auth

import (
	"fmt"
	"net/http"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
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

			existingUser, err := authStore.GetUserByEmail(r.Context(), req.Email)
			if err != nil {
				core.EncodeInternalError(w)
				fmt.Printf("%v\n", err)
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

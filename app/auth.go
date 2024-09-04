package app

import (
	"context"
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/andrew-hayworth22/rate-my-media/business"
)

type AppUser struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type PostUserRequest struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	DisplayName          string `json:"display_name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (a *App) PostUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req PostUserRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		ValidationError(w, business.ValidationError{
			Errors: map[string]string{
				"error": "Malformed JSON",
			},
		})
		return
	}

	email, err := mail.ParseAddress(req.Email)
	if err != nil {
		ValidationError(w, business.ValidationError{
			Errors: map[string]string{
				"email": "Invalid address",
			},
		})
		return
	}

	busReq := business.BusStoreUserRequest{
		Email:                *email,
		Name:                 req.Name,
		DisplayName:          req.DisplayName,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	busUsr, err := a.bus.CreateUser(ctx, busReq)
	if err != nil {
		vError, ok := err.(*business.ValidationError)
		if ok {
			ValidationError(w, *vError)
			return
		}
		InternalError(w, err)
		return
	}

	appUsr := AppUser{
		Id:          busUsr.Id,
		Email:       busUsr.Email.Address,
		Name:        busUsr.Name,
		DisplayName: busUsr.DisplayName,
	}
	Success(w, http.StatusOK, appUsr)
}

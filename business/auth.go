package business

import (
	"context"
	"net/mail"
	"strings"

	"github.com/andrew-hayworth22/rate-my-media/database"
)

type BusUser struct {
	Id           int
	Email        mail.Address
	Name         string
	DisplayName  string
	PasswordHash string
}

type BusStoreUserRequest struct {
	Email                mail.Address
	Name                 string
	DisplayName          string
	Password             string
	PasswordConfirmation string
}

func (b Business) CreateUser(ctx context.Context, req BusStoreUserRequest) (BusUser, error) {
	validationError := ValidationError{
		Errors: map[string]string{},
	}

	trimmedName := strings.Trim(req.Name, TRIM_SET)
	if len(trimmedName) == 0 {
		validationError.Errors["name"] = validText("name")
	}

	trimmedDisplayName := strings.Trim(req.DisplayName, TRIM_SET)
	if len(trimmedDisplayName) == 0 {
		validationError.Errors["display_name"] = validText("display name")
	}

	if len(req.Password) == 0 {
		validationError.Errors["password"] = validText("password")
	}

	if req.Password != req.PasswordConfirmation {
		validationError.Errors["password_confirmation"] = "These passwords do not match"
	}

	if len(validationError.Errors) > 0 {
		return BusUser{}, &validationError
	}

	dbReq := database.DbStoreUserRequest{
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Email:       req.Email.Address,
		Password:    req.Password,
	}
	dbUsr, err := b.db.CreateUser(ctx, dbReq)
	if err != nil {
		return BusUser{}, err
	}

	email, err := mail.ParseAddress(dbReq.Email)
	if err != nil {
		return BusUser{}, err
	}

	return BusUser{
		Id:           dbUsr.Id,
		Name:         dbUsr.Name,
		DisplayName:  dbUsr.DisplayName,
		Email:        *email,
		PasswordHash: dbUsr.PasswordHash,
	}, nil
}

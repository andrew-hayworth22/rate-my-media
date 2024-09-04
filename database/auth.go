package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type DbUser struct {
	Id           int
	Email        string
	Name         string
	DisplayName  string
	PasswordHash string
}

type DbStoreUserRequest struct {
	Id          int
	Email       string
	Name        string
	DisplayName string
	Password    string
}

func (db Database) CreateUser(ctx context.Context, req DbStoreUserRequest) (DbUser, error) {
	conn, err := db.Connect()
	if err != nil {
		return DbUser{}, err
	}
	defer conn.Close(ctx)

	conn.Begin(ctx)

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return DbUser{}, err
	}
	req.Password = string(encryptedPassword)

	sql := `
		insert into users (email, name, display_name, password)
		values(@email, @name, @display_name, @password)
		returning id
	`
	args := pgx.NamedArgs{
		"email":        req.Email,
		"name":         req.Name,
		"display_name": req.DisplayName,
		"password":     req.Password,
	}
	var id int
	if err = conn.QueryRow(ctx, sql, args).Scan(&id); err != nil {
		return DbUser{}, err
	}

	return DbUser{
		Id:           id,
		Email:        req.Email,
		Name:         req.Name,
		DisplayName:  req.DisplayName,
		PasswordHash: req.Password,
	}, nil
}

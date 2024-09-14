package auth

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type StorePg struct {
	dbUrl string
}

func NewStorePg() *StorePg {
	godotenv.Load()
	dbUrl := os.Getenv("DATABASE_URL")

	return &StorePg{
		dbUrl: dbUrl,
	}
}

func (asp *StorePg) Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), asp.dbUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func (asp *StorePg) StoreUser(ctx context.Context, req DbStoreUserRequest) (DbUser, error) {
	conn, err := asp.Connect()
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

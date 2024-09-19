package auth

import (
	"context"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type StorePg struct {
	dbUrl string
}

func NewAuthStorePg(dbUrl string) *StorePg {
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

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return DbUser{}, err
	}
	req.Password = string(encryptedPassword)

	transaction, err := conn.Begin(ctx)
	if err != nil {
		return DbUser{}, err
	}
	defer transaction.Rollback(ctx)

	sql := `
		insert into users (email, name, display_name, password, created_on)
		values(@email, @name, @display_name, @password, NOW())
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

	if err := transaction.Commit(ctx); err != nil {
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

func (asp *StorePg) GetUserByEmail(ctx context.Context, email string) (DbUser, error) {
	conn, err := asp.Connect()
	if err != nil {
		return DbUser{}, err
	}
	defer conn.Close(ctx)

	sql := `
		select id, email, name, display_name, password
		from users
		where email = $1;
	`
	rows, err := conn.Query(ctx, sql, email)
	if err != nil {
		return DbUser{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[DbUser])
	if err != nil {
		return DbUser{}, err
	}

	return user, nil
}

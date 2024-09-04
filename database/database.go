package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type DatabaseLayer interface {
	CreateUser(context.Context, DbStoreUserRequest) (DbUser, error)
}

type DatabaseConnection interface {
	Close()
}

type Database struct {
	dbUrl string
}

func NewDatabase() Database {
	godotenv.Load()
	dbUrl := os.Getenv("DATABASE_URL")

	return Database{
		dbUrl: dbUrl,
	}
}

func (db *Database) Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), db.dbUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

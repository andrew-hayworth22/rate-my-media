package media

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type MediaStore interface {
	DeleteMedia(ctx context.Context, id int) error
}

type PgMediaStore struct {
	dbUrl string
}

func NewMediaStorePg(dbUrl string) *PgMediaStore {
	return &PgMediaStore{
		dbUrl: dbUrl,
	}
}

func (asp *PgMediaStore) Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), asp.dbUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

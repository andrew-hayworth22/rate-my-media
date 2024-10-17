package movies

import (
	"context"

	"github.com/andrew-hayworth22/rate-my-media/database"
	"github.com/jackc/pgx/v5"
)

type MovieStore interface {
	StoreMovie(ctx context.Context, req DbStoreMovieRequest) (DbMovie, error)
	GetMovies(ctx context.Context) ([]DbMovie, error)
	GetMovieById(ctx context.Context, id int) (DbMovie, error)
	UpdateMovie(ctx context.Context, req DbUpdateMovieRequest) (DbMovie, error)
}

type PgMovieStore struct {
	dbUrl string
}

func NewMovieStorePg(dbUrl string) *PgMovieStore {
	return &PgMovieStore{
		dbUrl: dbUrl,
	}
}

func (asp *PgMovieStore) Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), asp.dbUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type DbMovie struct {
	database.DbMedia
	RuntimeMinutes int `db:"runtime_minutes"`
}

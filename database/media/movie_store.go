package media

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type MovieStore interface {
	StoreMovie(ctx context.Context, req DbStoreMovieRequest) (DbMovie, error)
	GetMovieById(ctx context.Context, id int) (DbMovie, error)
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

func (msp *PgMovieStore) StoreMovie(ctx context.Context, req DbStoreMovieRequest) (DbMovie, error) {
	conn, err := msp.Connect()
	if err != nil {
		return DbMovie{}, err
	}
	defer conn.Close(ctx)

	transaction, err := conn.Begin(ctx)
	if err != nil {
		return DbMovie{}, err
	}
	defer transaction.Rollback(ctx)

	sql := `
		insert into media (media_type_id, name, description, release_date)
		values (@media_type_id, @name, @description, @release_date)
		returning id;
	`
	args := pgx.NamedArgs{
		"media_type_id": MEDIA_TYPE_MOVIE.Id,
		"name":          req.Name,
		"description":   req.Description,
		"release_date":  req.ReleaseDate,
	}
	var id int
	if err = conn.QueryRow(ctx, sql, args).Scan(&id); err != nil {
		return DbMovie{}, err
	}

	sql = `
		insert into movies (id, runtime_minutes)
		values (@id, @runtime_minutes)
	`
	args = pgx.NamedArgs{
		"id":              id,
		"runtime_minutes": req.RuntimeMinutes,
	}
	if _, err := conn.Exec(ctx, sql, args); err != nil {
		return DbMovie{}, err
	}

	if err := transaction.Commit(ctx); err != nil {
		return DbMovie{}, err
	}

	return DbMovie{
		DbMedia: DbMedia{
			Id:          id,
			MediaType:   MEDIA_TYPE_MOVIE,
			Name:        req.Name,
			Description: req.Description,
			ReleaseDate: req.ReleaseDate,
		},
		RuntimeMinutes: req.RuntimeMinutes,
	}, nil
}

func (msp *PgMovieStore) GetMovieById(ctx context.Context, id int) (DbMovie, error) {
	conn, err := msp.Connect()
	if err != nil {
		return DbMovie{}, err
	}
	defer conn.Close(ctx)

	sql := `
		select media.id as id, media.name as name, media.description as description, media.release_date as release_date, movies.runtime_minutes as runtime_minutes
		from movies
		join media on movies.id = media.id
		where movies.id = $1;
	`
	rows, err := conn.Query(ctx, sql, id)
	if err != nil {
		return DbMovie{}, err
	}

	movie, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[DbMovie])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return DbMovie{}, nil
		}
		return DbMovie{}, err
	}

	movie.MediaType = MEDIA_TYPE_MOVIE

	return movie, nil
}

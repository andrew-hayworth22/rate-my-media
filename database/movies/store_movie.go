package movies

import (
	"context"
	"time"

	"github.com/andrew-hayworth22/rate-my-media/database"
	"github.com/jackc/pgx/v5"
)

type DbStoreMovieRequest struct {
	Name           string
	Description    string
	ReleaseDate    time.Time
	RuntimeMinutes int
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
		"media_type_id": database.MEDIA_TYPE_MOVIE.Id,
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
		DbMedia: database.DbMedia{
			Id:          id,
			MediaType:   database.MEDIA_TYPE_MOVIE,
			Name:        req.Name,
			Description: req.Description,
			ReleaseDate: req.ReleaseDate,
		},
		RuntimeMinutes: req.RuntimeMinutes,
	}, nil
}

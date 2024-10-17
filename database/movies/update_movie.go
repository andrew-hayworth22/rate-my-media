package movies

import (
	"context"
	"time"

	"github.com/andrew-hayworth22/rate-my-media/database"
	"github.com/jackc/pgx/v5"
)

type DbUpdateMovieRequest struct {
	Id             int
	Name           string
	Description    string
	ReleaseDate    time.Time
	RuntimeMinutes int
}

func (msp *PgMovieStore) UpdateMovie(ctx context.Context, req DbUpdateMovieRequest) (DbMovie, error) {
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
		update media
		set name = @name, description = @description, release_date = @release_date
		where id = @id;
	`
	args := pgx.NamedArgs{
		"id":           req.Id,
		"name":         req.Name,
		"description":  req.Description,
		"release_date": req.ReleaseDate,
	}

	if _, err := conn.Exec(ctx, sql, args); err != nil {
		return DbMovie{}, err
	}

	sql = `
		update movies
		set runtime_minutes = @runtime_minutes
		where id = @id;
	`
	args = pgx.NamedArgs{
		"id":              req.Id,
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
			Id:          req.Id,
			MediaType:   database.MEDIA_TYPE_MOVIE,
			Name:        req.Name,
			Description: req.Description,
			ReleaseDate: req.ReleaseDate,
		},
		RuntimeMinutes: req.RuntimeMinutes,
	}, nil
}

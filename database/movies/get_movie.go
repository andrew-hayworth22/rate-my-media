package movies

import (
	"context"
	"errors"

	"github.com/andrew-hayworth22/rate-my-media/database"
	"github.com/jackc/pgx/v5"
)

func (msp *PgMovieStore) GetMovies(ctx context.Context) ([]DbMovie, error) {
	conn, err := msp.Connect()
	if err != nil {
		return []DbMovie{}, err
	}
	defer conn.Close(ctx)

	sql := `
		select media.id as id, media.name as name, media.description as description, media.release_date as release_date, movies.runtime_minutes as runtime_minutes
		from movies
		join media on movies.id = media.id
		order by media.name;
	`
	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return []DbMovie{}, err
	}

	movies, err := pgx.CollectRows(rows, pgx.RowToStructByName[DbMovie])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []DbMovie{}, nil
		}
		return []DbMovie{}, err
	}

	return movies, nil
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

	movie.MediaType = database.MEDIA_TYPE_MOVIE

	return movie, nil
}

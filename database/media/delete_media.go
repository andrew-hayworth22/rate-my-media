package media

import (
	"context"
)

func (msp *PgMediaStore) DeleteMedia(ctx context.Context, id int) error {
	conn, err := msp.Connect()
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	sql := `
		delete from media
		where id = $1;
	`
	_, err = conn.Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}

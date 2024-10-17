package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"github.com/andrew-hayworth22/rate-my-media/app"
	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"github.com/andrew-hayworth22/rate-my-media/database/media"
	"github.com/andrew-hayworth22/rate-my-media/database/movies"
	"github.com/andrew-hayworth22/rate-my-media/migrate"
	"github.com/joho/godotenv"
)

func run(ctx context.Context, w io.Writer, getenv func(string) string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	dbUrl := getenv("DATABASE_URL")
	cfg := core.Config{
		JwtSecret: getenv("JWT_SECRET"),
	}

	handler := app.NewServer(cfg, auth.NewAuthStorePg(dbUrl), media.NewMediaStorePg(dbUrl), movies.NewMovieStorePg(dbUrl))

	port := getenv("PORT")
	server := http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	if _, err := w.Write([]byte("Running server on port: " + port + "\n")); err != nil {
		return err
	}
	return server.ListenAndServe()
}

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
		os.Exit(1)
	}

	migrateFlag := flag.Bool("migrate", false, "Run all outstanding migrations")
	migrateFreshFlag := flag.Bool("migrate-fresh", false, "Clear database and run all migrations")
	flag.Parse()
	if *migrateFlag {
		migrate.MigrateDB(ctx, os.Getenv, false)
	} else if *migrateFreshFlag {
		migrate.MigrateDB(ctx, os.Getenv, true)
	}

	if err := run(ctx, os.Stdout, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

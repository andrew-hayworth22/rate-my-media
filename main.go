package main

import (
	"context"
	"fmt"
	"github.com/andrew-hayworth22/rate-my-media/app"
	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"os/signal"
)

func run(ctx context.Context, w io.Writer, args []string, getenv func(string) string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg := core.Config{
		JwtSecret: getenv("JWT_SECRET"),
	}

	handler := app.NewServer(cfg, auth.NewAuthStorePg(getenv("DATABASE_URL")))

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

	if err := run(ctx, os.Stdout, os.Args, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

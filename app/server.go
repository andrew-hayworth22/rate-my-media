package app

import (
	"net/http"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"github.com/andrew-hayworth22/rate-my-media/database/media"
)

func NewServer(
	cfg core.Config,
	authStore auth.Store,
	movieStore media.MovieStore,
) http.Handler {
	mux := http.NewServeMux()

	AddRoutes(mux, cfg, authStore, movieStore)

	var handler http.Handler = mux

	return handler
}

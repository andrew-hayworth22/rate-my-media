package app

import (
	"net/http"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"github.com/andrew-hayworth22/rate-my-media/database/media"
	"github.com/andrew-hayworth22/rate-my-media/database/movies"
)

func NewServer(
	cfg core.Config,
	authStore auth.Store,
	mediaStore media.MediaStore,
	movieStore movies.MovieStore,
) http.Handler {
	mux := http.NewServeMux()

	AddRoutes(mux, cfg, authStore, movieStore, mediaStore)

	var handler http.Handler = mux

	return handler
}

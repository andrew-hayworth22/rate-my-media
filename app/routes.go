package app

import (
	"net/http"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	authHandlers "github.com/andrew-hayworth22/rate-my-media/app/handlers/auth"
	"github.com/andrew-hayworth22/rate-my-media/app/handlers/movies"
	authDb "github.com/andrew-hayworth22/rate-my-media/database/auth"
	"github.com/andrew-hayworth22/rate-my-media/database/media"
)

func AddRoutes(mux *http.ServeMux, cfg core.Config, authStore authDb.Store, movieStore media.MovieStore) {
	mux.Handle("/api/users", core.Post(authHandlers.HandlePostUser(authStore)))
	mux.Handle("/api/login", core.Post(authHandlers.HandleLogin(cfg, authStore)))

	mux.Handle("/api/movies", core.Post(movies.HandlePostMovie(movieStore)))
}

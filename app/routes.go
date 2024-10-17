package app

import (
	"net/http"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	authHandlers "github.com/andrew-hayworth22/rate-my-media/app/handlers/auth"
	moviesHandlers "github.com/andrew-hayworth22/rate-my-media/app/handlers/movies"
	authDb "github.com/andrew-hayworth22/rate-my-media/database/auth"
	moviesDb "github.com/andrew-hayworth22/rate-my-media/database/movies"
)

func AddRoutes(mux *http.ServeMux, cfg core.Config, authStore authDb.Store, movieStore moviesDb.MovieStore) {
	mux.Handle("POST /api/users", authHandlers.HandlePostUser(authStore))
	mux.Handle("POST /api/login", authHandlers.HandleLogin(cfg, authStore))

	mux.Handle("GET /api/movies", moviesHandlers.HandleGetMovies(movieStore))
	mux.Handle("GET /api/movies/{id}", moviesHandlers.HandleGetMovie(movieStore))
	mux.Handle("POST /api/movies", core.Authenticated(cfg, moviesHandlers.HandlePostMovie(movieStore)))
	mux.Handle("PUT /api/movies/{id}", core.Authenticated(cfg, moviesHandlers.HandlePutMovie(movieStore)))
}

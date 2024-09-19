package app

import (
	"github.com/andrew-hayworth22/rate-my-media/app/core"
	authHandlers "github.com/andrew-hayworth22/rate-my-media/app/handlers/auth"
	authDb "github.com/andrew-hayworth22/rate-my-media/database/auth"
	"net/http"
)

func AddRoutes(mux *http.ServeMux, cfg core.Config, authStore authDb.Store) {
	mux.Handle("/api/users", authHandlers.HandlePostUser(authStore))
	mux.Handle("/api/login", authHandlers.HandleLogin(cfg, authStore))
}

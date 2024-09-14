package app

import (
	"github.com/andrew-hayworth22/rate-my-media/app/core"
	authHandlers "github.com/andrew-hayworth22/rate-my-media/app/handlers/auth"
	authDb "github.com/andrew-hayworth22/rate-my-media/database/auth"
	"net/http"
)

func AddRoutes(mux *http.ServeMux, authStore authDb.Store, cfg core.Config) {
	mux.Handle("/api/login", authHandlers.HandlePostUser(authStore, cfg))
}

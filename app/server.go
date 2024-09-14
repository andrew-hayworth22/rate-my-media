package app

import (
	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/auth"
	"net/http"
)

func NewServer(
	cfg core.Config,
	authStore *auth.Store,
) *http.Handler {

}

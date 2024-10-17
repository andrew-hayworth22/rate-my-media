package core

import (
	"context"
	"net/http"
)

func Authenticated(cfg Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tok := r.Header.Get("Authorization")
		if len(tok) == 0 {
			EncodeUnauthorized(w)
			return
		}

		jwt, err := DecodeJWT(cfg, tok)
		if err != nil {
			EncodeUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), "jwt", jwt)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

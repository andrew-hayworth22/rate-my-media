package movies

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/movies"
)

type PostMovieRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	ReleaseDate    string `json:"release_date"`
	RuntimeMinutes int    `json:"runtime_minutes"`
}

func (req PostMovieRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = map[string]string{}

	trimmedName := strings.TrimSpace(req.Name)
	if len(trimmedName) == 0 {
		problems["name"] = "Please enter a valid name"
	}

	trimmedReleaseDate := strings.TrimSpace(req.ReleaseDate)
	if len(trimmedReleaseDate) == 0 {
		problems["release_date"] = "Please enter a date"
	} else if _, err := time.Parse(time.RFC3339, trimmedReleaseDate); err != nil {
		problems["release_date"] = "Malformed date"
	}

	if req.RuntimeMinutes < 1 {
		problems["runtime_minutes"] = "Please enter a valid runtime"
	}

	return problems
}

func HandlePostMovie(movieStore movies.MovieStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var req PostMovieRequest
			req, problems, err := core.DecodeValid(r, req)
			if err != nil {
				core.EncodeInternalError(w)
				return
			}
			if len(problems) > 0 {
				core.EncodeValidationError(w, problems)
				return
			}

			parsedDate, err := time.Parse(time.RFC3339, req.ReleaseDate)
			if err != nil {
				core.EncodeInternalError(w)
				return
			}

			dbMovie, err := movieStore.StoreMovie(r.Context(), movies.DbStoreMovieRequest{
				Name:           req.Name,
				Description:    req.Description,
				ReleaseDate:    parsedDate,
				RuntimeMinutes: req.RuntimeMinutes,
			})
			if err != nil {
				fmt.Println(err)
				core.EncodeInternalError(w)
				return
			}

			appMovie := AppMovie{
				Id:             dbMovie.Id,
				Name:           dbMovie.Name,
				Description:    dbMovie.Description,
				ReleaseDate:    dbMovie.ReleaseDate,
				RuntimeMinutes: dbMovie.RuntimeMinutes,
			}
			core.Encode(w, http.StatusOK, appMovie)
		})

}

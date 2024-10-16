package movies

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/media"
)

func HandleGetMovie(movieStore media.MovieStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			idString := r.PathValue("id")
			id, err := strconv.Atoi(idString)
			if err != nil {
				core.EncodeNotFound(w)
				return
			}

			dbMovie, err := movieStore.GetMovieById(r.Context(), id)
			if err != nil {
				fmt.Printf("%v", err)
				core.EncodeInternalError(w)
				return
			}
			if dbMovie.Id == 0 {
				core.EncodeNotFound(w)
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

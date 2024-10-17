package media

import (
	"net/http"
	"strconv"

	"github.com/andrew-hayworth22/rate-my-media/app/core"
	"github.com/andrew-hayworth22/rate-my-media/database/media"
)

func HandleDeleteMedia(mediaStore media.MediaStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			idString := r.PathValue("id")
			id, err := strconv.Atoi(idString)
			if err != nil {
				core.EncodeNotFound(w)
				return
			}

			err = mediaStore.DeleteMedia(r.Context(), id)
			if err != nil {
				core.EncodeInternalError(w)
				return
			}

			core.Encode(w, http.StatusOK, "")
		})

}

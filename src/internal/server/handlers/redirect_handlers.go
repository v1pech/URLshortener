package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRedirect(logger *slog.Logger, db dbGetter) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			sendResponse(w, http.StatusBadRequest, "", fmt.Errorf("alias is empty"))
			return
		}
		URL, errNum, err := db.GetUrl(alias)
		switch errNum {
		default:
			http.Redirect(w, r, URL, http.StatusFound)
		case -1:
			sendResponse(w, http.StatusInternalServerError, "", fmt.Errorf("database error"))
			logger.Error(err.Error())
		case 2:
			sendResponse(w, http.StatusNotFound, "", fmt.Errorf("alias does not exist"))
		}

	}
	return http.HandlerFunc(fn)
}

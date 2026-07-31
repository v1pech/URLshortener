package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type dbGetter interface {
	GetUrl(alias string) (string, int8, error)
}

type dbSaver interface {
	SaveURL(url, alias string) (int8, error)
}

type dbDeleter interface {
	DeleteURL(alias string) (int8, error)
}

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}

// GET /alias/{alias}
func NewGetAlias(logger *slog.Logger, db dbGetter) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			sendResponse(w, http.StatusBadRequest, "", fmt.Errorf("alias is empty"))
			return
		}
		link, errNum, err := db.GetUrl(alias)
		switch errNum {
		default:
			sendResponse(w, http.StatusOK, link, nil)
		case -1:
			sendResponse(w, http.StatusInternalServerError, "", fmt.Errorf("database error"))
			logger.Error(err.Error())
			return
		case 2:
			sendResponse(w, http.StatusNotFound, "", fmt.Errorf("alias does not exist"))
		}
	}

	return http.HandlerFunc(fn)
}

// POST /alias
func NewPostAlias(logger *slog.Logger, db dbSaver) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, "", fmt.Errorf("can not read"))
			logger.Error(err.Error())
			return
		}
		var req Request
		err = json.Unmarshal(body, &req)
		if err != nil {
			sendResponse(w, http.StatusBadRequest, "", fmt.Errorf("can not unmarshal"))
			logger.Error(err.Error())
			return
		}
		errNum, err := db.SaveURL(req.URL, req.Alias)
		switch errNum {
		default:
			sendResponse(w, http.StatusOK, "", nil)
		case -1:
			logger.Error(err.Error())
			sendResponse(w, http.StatusInternalServerError, "", fmt.Errorf("can not save"))
		case 1:
			logger.Debug(err.Error())
			sendResponse(w, http.StatusBadRequest, "", fmt.Errorf("alias already exists"))
		}
	}
	return http.HandlerFunc(fn)
}

// DELETE /alias/{alias}
func NewDeleteAlias(logger *slog.Logger, db dbDeleter) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			sendResponse(w, http.StatusBadRequest, "", fmt.Errorf("alias is empty"))
			return
		}
		errNum, err := db.DeleteURL(alias)
		switch errNum {
		default:
			sendResponse(w, http.StatusOK, "", nil)
			return
		case -1:
			logger.Error(err.Error())
			sendResponse(w, http.StatusInternalServerError, "", fmt.Errorf("can not delete"))
			return
		case 2:
			sendResponse(w, http.StatusNotFound, "", fmt.Errorf("alias does not exist"))
			return
		}
	}
	return http.HandlerFunc(fn)
}

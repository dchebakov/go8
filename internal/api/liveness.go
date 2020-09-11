package api

import (
	"net/http"

	"github.com/go-chi/render"
)

func (a API) HandleLive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, nil)
		return
	}
}

func (a API) HandleReady() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.books.Ping()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, nil)
		return
	}
}

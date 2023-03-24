package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func router() http.Handler {
	r := chi.NewRouter()

	r.Get("/estp", getEstpAuctions)

	return r
}

func Start() {
	r := router()
	http.ListenAndServe(":8080", r)
}

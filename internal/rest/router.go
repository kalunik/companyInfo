package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Heartbeat("/ping"))

	r.Get("/{taxId}", TaxHandler)

	return r
}

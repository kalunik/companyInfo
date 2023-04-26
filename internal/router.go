package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kalunik/companyInfo/internal/handler"
)

func NewRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{taxId}", handler.TaxHandler)

	return r
}

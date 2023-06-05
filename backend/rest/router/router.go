package router

import (
	ownmiddleware "cryptotracker/rest/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(ownmiddleware.Cors)
	r.Use(ownmiddleware.Hearthbeat)
}

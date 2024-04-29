package main

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"isustud.com/m/storage/mariadb"
)

func addRoutes(
	r *chi.Mux,
	logger *slog.Logger,
	storage *mariadb.Storage,
) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/tags", TagsGetHandler(logger))
		r.Post("/tags", TagsPostHandler(logger, storage))
		r.Get("/apikeys/{login}", KeyGetHandler(logger, storage))
	})
}

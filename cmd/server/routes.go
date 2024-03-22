package main

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func addRoutes(
	r *chi.Mux,
	logger *slog.Logger,
) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/tags", TagsGetHandler())
	})
}

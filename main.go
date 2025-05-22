package main

import (
	"log/slog"
	"net/http"
	"template/handler"
	"template/ui"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(time.Second * 30))
	r.Use(middleware.NoCache)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/admin/count", handler.AdminCount)
		r.Post("/admin", handler.CreateAdmin)
	})

	r.Group(func(r chi.Router) {
	})

	r.Group(func(r chi.Router) {
	})

	r.Handle("/*", http.FileServerFS(ui.SubFS()))

	s := http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	slog.Info("http://localhost:4000")
	slog.Info("http://localhost:4001 (hot refresh)")

	s.ListenAndServe()
}

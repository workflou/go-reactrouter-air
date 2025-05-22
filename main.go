package main

import (
	"log/slog"
	"net/http"
	"template/ui"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

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

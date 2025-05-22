package config

import (
	"flag"

	"golang.org/x/exp/slog"
)

func init() {
	flag.StringVar(&Dsn, "dsn", "db.sqlite", "database dsn")

	flag.Parse()

	slog.Info("[config] config loaded")
}

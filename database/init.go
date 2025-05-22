package database

import (
	"database/sql"
	"log/slog"
	"template/config"
	"template/schema"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open("sqlite", config.Dsn)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	goose.SetDialect("sqlite3")
	goose.SetBaseFS(schema.FS)

	if err = goose.Up(DB, "."); err != nil {
		panic(err)
	}

	slog.Info("[database] database initialized")
}

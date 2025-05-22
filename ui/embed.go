package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var FS embed.FS

func SubFS() fs.FS {
	sub, err := fs.Sub(FS, "dist")

	if err != nil {
		panic(err)
	}

	return sub
}

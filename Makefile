.PHONY: watch
watch:
	go tool air -c .air.toml

.PHONY: build
build: sqlc ui
	go build -o ./tmp/main .

.PHONY: ui
ui:
	npm run build --prefix ./ui

.PHONY: create-migration
create-migration:
	GOOSE_DRIVER=sqlite3 GOOSE_MIGRATION_DIR=./schema go tool goose -s create migration sql

.PHONY: sqlc
sqlc:
	go tool sqlc generate
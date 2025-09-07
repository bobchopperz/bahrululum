.PHONY: help db-create db-drop db-recreate migrate migrate-up migrate-down migrate-status migrate-create migrate-reset build run

help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/migrate cmd/migrations/main.go
	go build -o bin/usermgmt cmd/usermgmt/main.go

run:
	go run cmd/server/main.go

db-create:
	go run cmd/migrations/main.go create-db

db-drop:
	go run cmd/migrations/main.go drop-db

db-recreate:
	make db-drop
	make db-create
	make migrate-up

migrate-up:
	go run cmd/migrations/main.go up

migrate-down:
	go run cmd/migrations/main.go down

migrate-status:
	go run cmd/migrations/main.go status

migrate-reset:
	go run cmd/migrations/main.go reset

migrate-create:
	@if [ -z "$(name)" ]; then echo "Usage: make migrate-create name=migration_name"; exit 1; fi
	go run cmd/migrations/main.go create $(name) sql

db-setup:
	make db-create
	make migrate-up

test:
	go test ./...

clean:
	rm -rf bin/

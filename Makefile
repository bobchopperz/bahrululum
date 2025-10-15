.PHONY: help db-create db-drop db-recreate db-connect db-shell migrate migrate-up migrate-down migrate-status migrate-create migrate-reset seed seed-clean build run

help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/migrate cmd/migrations/main.go
	go build -o bin/usermgmt cmd/usermgmt/main.go
	go build -o bin/seeder cmd/seeder/main.go

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

db-connect: ## Connect to the database using psql
	@DB_HOST=$$(grep -A 5 "^database:" configs/config.yaml | grep "host:" | awk '{print $$2}' | tr -d '"'); \
	DB_PORT=$$(grep -A 5 "^database:" configs/config.yaml | grep "port:" | awk '{print $$2}' | tr -d '"'); \
	DB_USER=$$(grep -A 5 "^database:" configs/config.yaml | grep "user:" | awk '{print $$2}' | tr -d '"'); \
	DB_NAME=$$(grep -A 5 "^database:" configs/config.yaml | grep "name:" | awk '{print $$2}' | tr -d '"'); \
	echo "Connecting to PostgreSQL database '$$DB_NAME' at $$DB_HOST:$$DB_PORT as user '$$DB_USER'..."; \
	PGPASSWORD=$$(grep -A 5 "^database:" configs/config.yaml | grep "password:" | awk '{print $$2}' | tr -d '"') \
	psql -h $$DB_HOST -p $$DB_PORT -U $$DB_USER -d $$DB_NAME

db-shell: db-connect ## Alias for db-connect

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

seed: ## Seed database with sample data
	go run cmd/seeder/main.go

seed-courses: ## Seed courses only
	go run cmd/seeder/main.go -courses

seed-chapters: ## Seed course chapters only
	go run cmd/seeder/main.go -chapters

seed-contents: ## Seed course contents only
	go run cmd/seeder/main.go -contents

seed-clean: ## Clean all seeded data and reseed
	go run cmd/seeder/main.go -clean -all

test:
	go test ./...

clean:
	rm -rf bin/

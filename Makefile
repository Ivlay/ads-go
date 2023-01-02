include .env
.PHONY:
.SILENT:

build:
	go build -o ./.bin/ads-go cmd/ads-go/main.go

run: build
			./.bin/ads-go

migrateup:
	migrate -path internal/db/migrations -database "postgres://postgres:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "postgres://postgres:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down

migrateforce:
	migrate -path internal/db/migrations -database "postgres://postgres:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" force 1
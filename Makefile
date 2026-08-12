.PHONY: dev build run migrate

-include .env
export

dev:
	air

build:
	go build -o ./build/ ./cmd/main.go

run:
	go run ./cmd/main.go

migrate:
	go run ./cmd/migration/main.go

%:
	@:
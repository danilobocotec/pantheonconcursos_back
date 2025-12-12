.PHONY: help install build run test clean fmt lint docker-up docker-down

help:
	@echo "Available commands:"
	@echo "  make install      - Install dependencies"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make fmt          - Format code"
	@echo "  make lint         - Run linter"
	@echo "  make dev          - Run in development mode"

install:
	go mod download
	go mod tidy

build:
	go build -o bin/api cmd/api/main.go

run: build
	./bin/api

dev:
	go run cmd/api/main.go

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	rm -rf bin/
	rm -f coverage.out

fmt:
	go fmt ./...

lint:
	golangci-lint run

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

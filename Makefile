include .env

APP_NAME=go-fridge
MAIN=cmd/server/main.go
DOCS_DIR=internal/docs

tidy:
	go mod tidy

run:
	docker compose up --build

debug:
	APP_ENV=debug docker compose up --build

stop:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f app

swag:
	go run github.com/swaggo/swag/cmd/swag@v1.8.12 init -g ${MAIN} -o ${DOCS_DIR}

.PHONY: tidy run down swag
POSTGRES_HOST ?= localhost
POSTGRES_PORT ?= 5432
POSTGRES_USER ?= shop
POSTGRES_PASSWORD ?= shop
POSTGRES_DB ?= shop
DSN ?= user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) sslmode=disable

.PHONY: test-env-up
test-env-up:
	docker-compose up -d

.PHONY: migration-create
migration-create:
	goose -dir migrations create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir migrations postgres "$(DSN)" up

.PHONY: migration-down
migration-down:
	goose -dir migrations postgres "$(DSN)" down

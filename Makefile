Migration_Folder=db/migrations/

-include .env

export GOOSE_DRIVER ?= mysql
export GOOSE_DBSTRING ?= $(DBUSER)@$(NET)($(ADDR))/$(DBNAME)?parseTime=true

.PHONY: migrate-create migrate-up migrate-down

migrate-create:
	goose -dir $(Migration_Folder) create $(name) sql

migrate-up:
	goose -dir $(Migration_Folder) up

migrate-down:
	goose -dir $(Migration_Folder) down
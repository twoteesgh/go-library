include .env
mysql_url="mysql://${DB_USER}:${DB_PASSWORD}@(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true"

.PHONY: all migrate-create migrate-up

all:
	$(info Please provide an argument)

migrate-version:
	@migrate -path migrations -database $(mysql_url) version

migrate-create:
	$(if $n,, $(error Please provide a migration name with n=))
	@migrate create -ext sql -dir migrations $n

migrate-up:
	@migrate -path migrations -database $(mysql_url) up $n

migrate-down:
	@migrate -path migrations -database $(mysql_url) down $n
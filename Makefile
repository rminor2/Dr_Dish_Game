include .env

MIGRATIONS=$(PWD)/db/migrations

test:
	echo $(DB_URL)

gen:
	sqlc generate

drop:
	migrate -source file://$(MIGRATIONS) -database $(DB_URL) drop

up:
	migrate -source file://$(MIGRATIONS) -database $(DB_URL) up

enterdb:
	psql $(DB_URL)


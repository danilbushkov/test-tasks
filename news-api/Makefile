-include .env

all: run

run:
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down

env:
	[ ! -f .env ] && cp .env.example .env; true

db:
	psql -h localhost -p 5432 -U ${POSTGRES_USER}

clear:
	rm -r database/pgdata

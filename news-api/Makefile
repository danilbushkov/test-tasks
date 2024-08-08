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
	PGPASSWORD=${POSTGRES_PASSWORD} psql \
		-h localhost \
		-p 5432 \
		-U ${POSTGRES_USER} \
		${POSTGRES_DB}

clear:
	rm -r database/pgdata

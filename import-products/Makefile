all: run


run:
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down

clear:
	rm -rf ./storage/app/public/pictures
	rm -rf ./database/pgdata

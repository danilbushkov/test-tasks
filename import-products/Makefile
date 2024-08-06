all: run



run:
	docker compose up nginx -d

stop:
	docker compose stop

down:
	docker compose down

migrate:
	docker compose run artisan migrate


clear:
	rm -rf ./storage/app/public/pictures
	rm -rf ./database/pgdata
	rm -f ./public/storage

rights:
	-chmod -R o+w ./storage
	-chmod -R o+x .

mkdirs:
	mkdir ./storage/app/public/pictures

build: clear mkdirs rights
	cp .env.example .env
	docker compose run composer install
	docker compose run artisan storage:link
	docker compose run artisan key:generate



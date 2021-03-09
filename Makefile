.SILENT:

docker-dependencies = Dockerfile docker-compose.yml wait-for-postgres.sh

build: $(docker-dependencies)
	docker-compose build to-do-app

up: $(docker-dependencies)
	docker-compose up to-do-app

clean: $(docker-dependencies)
	docker-compose stop db
	docker-compose stop to-do-app
	docker-compose rm db to-do-app

migrate-up: ./schema
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

migrate-down: ./schema
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' down
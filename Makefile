.SILENT:

test:
	go test ./pkg/handler/
	go test ./pkg/repository/

build:
	docker-compose build to-do-app

up:
	docker-compose up to-do-app

down:
	docker-compose stop down

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' down

swag:
	swag init -g cmd/main.go

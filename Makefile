DB_URL=postgresql://root:root@localhost:5432/task_manager?sslmode=disable

startdb:
	@echo "Starting Postgres..."
	brew services start postgresql # for mac
	docker compose up -d db # for other platform with docker installed
	psql postgres
	CREATE USER root WITH SUPERUSER PASSWORD 'root';
	CREATE DATABASE task_manager;

stopdb:
	docker compose down -v --remove-orphans --rmi all

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/litmus-zhang/task_manager/db/sqlc Store

.PHONY: startdb stopdb migrateup migrateup1 migratedown migratedown1 new_migration sqlc test server mock
.PHONY: sqlc serve createdb dropdb dbup dbdown postgres

sqlc:
	sqlc generate

serve:
	go run main.go

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root kins_db

dropdb:
	docker exec -it postgres12 dropdb kins_db

dbup:
	migrate -path database/migration -database "postgresql://root:123456@localhost:5432/kins_db?sslmode=disable" -verbose up

dbdown:
	migrate -path database/migration -database "postgresql://root:123456@localhost:5432/kins_db?sslmode=disable" -verbose down

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:12-alpine

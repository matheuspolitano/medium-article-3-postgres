postgres:
	docker run --name postgres-example -p 5432:5432 -e POSTGRES_PASSWORD=MP@TEST123 -e POSTGRES_USER=root -d postgres:12-alpine

createdb:
	docker exec -it postgres-example createdb --username=root --owner=root simple_table

dropdb:
	docker exec -it postgres-example dropdb  simple_table

migrationup:
	migrate -database "postgresql://root:MP@TEST123@localhost:5432/simple_table?sslmode=disable" -path db/migrations up

migrationdown:
	migrate -database "postgresql://root:MP@TEST123@localhost:5432/simple_table?sslmode=disable" -path db/migrations down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrationup migrationdown sqlc

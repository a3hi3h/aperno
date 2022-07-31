DB_URL=postgresql://root:secretpwd@localhost:5432/aperno?sslmode=disable

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpwd -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root aperno

dropdb:
	docker exec -it postgres dropdb aperno

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

sqlcdocker:
	docker pull kjconroy/sqlc
	docker run --rm -v "%cd%:/aperno" -w /aperno kjconroy/sqlc generate

test:
	go test -v -cover ./...

main:
	go run main.go

.PHONY:  postgres createdb dropdb migrateup migratedown sqlc sqlcdocker test main

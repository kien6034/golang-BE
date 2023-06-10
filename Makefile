DB_URL=postgresql://root:kien6034@localhost:5432/simple_bank?sslmode=disable
.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis


postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=kien6034 -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

startdb:
	docker start postgres15

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down


sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/kien6034/golang-BE/db/sqlc Store

server:
	go run main.go

test:
	go test -v -cover ./...
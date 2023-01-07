DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
postgres-restart:
	docker restart postgres
postgres-stop:
	docker stop postgres

create-db:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
drop-db:
	docker exec -it postgres dropdb simple_bank

migrate-up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrate-down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sql-c:
	sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc server test

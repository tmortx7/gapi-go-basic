DB_URL=postgresql://root:secret@localhost:5432/dbdata?sslmode=disable

network:
	docker network create mynetwork

postgres:
	docker run --name postgres --network mynetwork -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root dbdata

dropdb:
	docker exec -it postgres dropdb dbdata

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up


migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down



db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 db_docs db_schema sqlc test server mock proto evans
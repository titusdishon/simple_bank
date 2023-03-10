DB_URL=postgresql://root:geek36873@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name postgres15  --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=geek36873 -d postgres:15-alpine

createDb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropDb:
	docker exec -t postgres15 dropdb  simple_bank

migrateUp:
	migrate -path db/migrations -database  -verbose up

migrateUpTest:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrateUp1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up 1

migrateDown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

migrateDown1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./... -coverprofile coverage.out

coverage:
	chmod u+x ./coverage.sh && ./coverage.sh

server:
	go run main.go
	
mock:
	mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/titusdishon/simple_bank/db/sqlc Store 

dbdocs:
	dbdocs build docs/db.dbml
dbschema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml
proto:
	rm -f pb/*go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto
evans:
	 evans --host localhost --port 8090 -r repl
.PHONY: createDb, postgres, dropDb, migrateUp, migrateDown, migrateUp1, migrateDown1 ,sqlc, server, mock, coverage,dbdocs,dbschema, proto, evans


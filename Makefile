postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=geek36873 -d postgres:15-alpine

createDb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropDb:
	docker exec -t postgres15 dropdb  simple_bank

migrateUp:
	migrate -path db/migrations -database "postgresql://root:geek36873@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migrations -database "postgresql://root:geek36873@localhost:5432/simple_bank?sslmode=disable" -verbose down

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

.PHONY: createDb, postgres, dropDb, migrateUp, migrateDown ,sqlc, server, mock, coverage
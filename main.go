package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/titusdishon/simple_bank/api"
	db "github.com/titusdishon/simple_bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:geek36873@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(" Cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewSever(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal(" Cannot start server: ", err)
	}
}

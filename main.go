package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/osmanok/simple-bank-backend/api"
	db "github.com/osmanok/simple-bank-backend/db/sqlc"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgres://osmanokuyan:Gok898@localhost:5432/simplebank?sslmode=disable"
	serverAdress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect the db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}

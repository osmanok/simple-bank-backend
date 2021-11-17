package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/osmanok/simple-bank-backend/api"
	db "github.com/osmanok/simple-bank-backend/db/sqlc"
	"github.com/osmanok/simple-bank-backend/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect the db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}

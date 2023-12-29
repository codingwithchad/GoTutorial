package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	// Create database connection
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: %w", err)
	}

	// Create a store
	store := db.NewStore(conn)

	// Create a new server
	server := api.NewServer(store)

	// Start the server
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: %w", err)
	}
}

package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/codingwithchad/simplebank/api"
	db "github.com/codingwithchad/simplebank/db/sqlc"
	"github.com/codingwithchad/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	// Create database connection
	conn, err := sql.Open(config.DBDriver, config.DBSources)
	if err != nil {
		log.Fatal("cannot connect to db: %w", err)
	}

	// Create a store
	store := db.NewStore(conn)

	// Create a new server
	server := api.NewServer(store)

	// Start the server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: %w", err)
	}
}

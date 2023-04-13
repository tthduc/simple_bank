package main

import (
	"database/sql"
	"example.com/m/api"
	"example.com/m/internal/db"
	"example.com/m/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig("app.yaml")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

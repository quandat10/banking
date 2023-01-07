package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/quandat10/banking/api"
	db "github.com/quandat10/banking/db/sqlc"
	"github.com/quandat10/banking/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}
